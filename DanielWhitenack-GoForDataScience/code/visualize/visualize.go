// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This example demonstrates how to create histograms with gonum/plot.  The program
// saves two histograms.  The first is a histogram of github star values per Go repo
// (see getrepos.go for information on how this data is retrieved), and the second
// is a histogram of the log of non-zero star values.
package main

import (
	"bytes"
	"encoding/csv"
	"log"
	"math"
	"strconv"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
	"github.com/pachyderm/pachyderm/src/client"
	"github.com/pkg/errors"
)

func main() {

	// Prepare the dataset in pachyderm for plotting.
	v, vl, err := prepareStarData("repodata")
	if err != nil {
		log.Fatal(err)
	}

	// Create and save the histogram plots.
	if err := makePlots(v, vl); err != nil {
		log.Fatal(err)
	}
}

// prepareStartData translates the previously stored data set into values for gonum/plot.
func prepareStarData(dataSet string) (plotter.Values, plotter.Values, error) {

	// Get the data set we stored in pachyderm.
	data, err := getDataSet(dataSet, "master", "godata")
	if err != nil {
		return nil, nil, errors.Wrap(err, "Could not get data from pachyderm")
	}

	// Extract the records from the data.
	reader := csv.NewReader(bytes.NewReader(data.Bytes()))
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, errors.Wrap(err, "Could not read in data records.")
	}

	v := make(plotter.Values, len(records))
	var vl plotter.Values

	// Loop over each record adding the star field into the plotting values.
	for i, each := range records {
		value, err := strconv.ParseFloat(each[5], 64)
		if err != nil {
			return nil, nil, errors.Wrap(err, "Could not convert value to float")
		}

		v[i] = value
		if value != 0.0 {
			vl = append(vl, math.Log(value))
		}
	}

	return v, vl, nil
}

// getDataSet gets a previously stored dataset from pachyderm data versioning.
func getDataSet(dataSet, branch, repoName string) (bytes.Buffer, error) {

	// Open a connection to pachyderm running on localhost.
	c, err := client.NewFromAddress("localhost:30650")
	if err != nil {
		return bytes.Buffer{}, errors.Wrap(err, "Could not connect to Pachyderm")
	}

	// Read the latest commit of filename to the given repoName.
	var buffer bytes.Buffer
	if err := c.GetFile(repoName, branch, dataSet, 0, 0, "", nil, &buffer); err != nil {
		return buffer, errors.Wrap(err, "Could not retrieve pachyderm file")
	}

	return buffer, nil
}

// makePlots creates and saves both histogram plots.
func makePlots(v, vl plotter.Values) error {

	// Create a new plot and set the title.
	p1, err := plot.New()
	if err != nil {
		return errors.Wrap(err, "Could not generate a new plot")
	}
	p1.Title.Text = "Histogram of Github Stars"

	// Create a second plot and set the title.
	p2, err := plot.New()
	if err != nil {
		return errors.Wrap(err, "Could not generate a new plot")
	}
	p2.Title.Text = "Histogram of log(Github Stars)"

	// Create a histogram and add it to plot 1.
	h1, err := plotter.NewHist(v, 16)
	if err != nil {
		return errors.Wrap(err, "Could not create histogram")
	}
	p1.Add(h1)

	// Create a histogram and add it to plot 2.
	h2, err := plotter.NewHist(vl, 16)
	if err != nil {
		return errors.Wrap(err, "Could not create histogram")
	}
	p2.Add(h2)

	// Save plot 1's histogram in the current directory.
	if err := p1.Save(4*vg.Inch, 4*vg.Inch, "hist1.png"); err != nil {
		return errors.Wrap(err, "Could not save plot")
	}

	// Save plot 2's histogram in the current directory.
	if err := p2.Save(4*vg.Inch, 4*vg.Inch, "hist2.png"); err != nil {
		return errors.Wrap(err, "Could not save plot")
	}

	return nil
}
