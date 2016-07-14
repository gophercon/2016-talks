// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This example demonstrates how to train a regression model in Go.  The example
// also prints out formatted results and saves two plot: (i) a plot of the raw input
// data, and (ii) a plot of the trained function overlaid on the raw input data.
// The input data is data about Go github repositories gathered via getrepos.go.
package main

import (
	"bytes"
	"encoding/csv"
	"log"
	"sort"
	"time"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
	"github.com/pachyderm/pachyderm/src/client"
	"github.com/pkg/errors"
)

func main() {

	// Aggregate the counts of created repos per day over all days.
	counts, err := prepareCountData("repodata")
	if err != nil {
		log.Fatal(err)
	}

	// Create and save the plot showing the time series of
	// observed daily created repo counts.
	xys := preparePlotData(counts)
	if err = makePlots(xys); err != nil {
		log.Fatal(err)
	}
}

// prepareCountData transforms the dataset into a series of
// daily count values for plotting.
func prepareCountData(dataSet string) ([][]int, error) {

	// Create a map to store the daily counts of created repos.
	countMap := make(map[int]int)

	// Get the data set we stored in pachyderm.
	data, err := getDataSet(dataSet, "master", "godata")
	if err != nil {
		return [][]int{}, errors.Wrap(err, "Could not get data from pachyderm")
	}

	// Extract the records from the data.
	reader := csv.NewReader(bytes.NewReader(data.Bytes()))
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		return [][]int{}, errors.Wrap(err, "Could not read in data records.")
	}

	// Create a map of daily created repos where the keys are the days and
	// the values are the counts of created repos on that day.
	startTime := time.Date(2013, time.January, 1, 0, 0, 0, 0, time.UTC)
	layout := "2006-01-02 15:04:05"
	for _, each := range records {
		t, err := time.Parse(layout, each[2][0:19])
		if err != nil {
			return [][]int{}, errors.Wrap(err, "Could not parse timestamps")
		}
		interval := int(t.Sub(startTime).Hours() / 24.0)
		countMap[interval]++
	}

	// Sort the day values which is required for plotting.
	var keys []int
	for k := range countMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var sortedCounts [][]int
	for _, k := range keys {
		sortedCounts = append(sortedCounts, []int{k, countMap[k]})
	}

	return sortedCounts, nil
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

// preparePlotData prepares the raw input data for plotting.
func preparePlotData(counts [][]int) plotter.XYs {
	pts := make(plotter.XYs, len(counts))
	var i int

	for _, count := range counts {
		pts[i].X = float64(count[0])
		pts[i].Y = float64(count[1])
		i++
	}

	return pts
}

// makePlots creates and saves the first of our plots showing the raw input data.
func makePlots(xys plotter.XYs) error {

	// Create a new plot.
	p, err := plot.New()
	if err != nil {
		return errors.Wrap(err, "Could not create plot object")
	}

	// Label the new plot.
	p.Title.Text = "Daily Counts of Go Repos Created"
	p.X.Label.Text = "Days from Jan. 1, 2013"
	p.Y.Label.Text = "Count"

	// Add the prepared points to the plot.
	if err = plotutil.AddLinePoints(p, "Counts", xys); err != nil {
		return errors.Wrap(err, "Could not add lines to plot")
	}

	// Save the plot to a PNG file.
	if err := p.Save(7*vg.Inch, 4*vg.Inch, "countseries.png"); err != nil {
		return errors.Wrap(err, "Could not output plot")
	}

	return nil
}
