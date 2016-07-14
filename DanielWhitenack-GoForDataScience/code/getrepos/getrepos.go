// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This example demonstrates how to scrape data about Go github repositories
// using the github.com/google/go-github/github package.
package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/google/go-github/github"
	"github.com/pkg/errors"
)

func main() {

	// Capture the date from which we will start scraping github data.
	t1 := time.Date(2013, time.January, 1, 0, 0, 0, 0, time.UTC)

	// Pull the github data from t1 to now.
	if err := queryFromStartTime(t1); err != nil {
		log.Fatal(err)
	}
}

// queryFromStartTime queries github for all 2 day time ranges of repo create
// dates from a start time until now.
func queryFromStartTime(t1 time.Time) error {

	// Create a new github client.
	client := github.NewClient(nil)

	// Loop over 2 days periods from the start time until now such that
	// the number of github repos we return is less than the 10 pages X
	// 100 results limit of the github API.
	for t1.Unix() < time.Now().Unix() {

		// Create the github query string.
		t2 := t1.Add(time.Hour * 24 * 2)
		tString := fmt.Sprintf("\"%d-%02d-%02d .. %d-%02d-%02d\"",
			t1.Year(), t1.Month(), t1.Day(),
			t2.Year(), t2.Month(), t2.Day())
		query := fmt.Sprintf("language:Go created:" + tString)

		// Query github with our new query.
		err := clientQuery(client, query)
		if err != nil {
			errors.Wrap(err, "Could not search Github repos")
		}

		// Increment the start time of the 2 day chunk.
		t1 = t1.Add(time.Hour * 24 * 2)
	}

	return nil
}

// clientQuery executes github queries and searches over all pages of a result
// set parsing results.
func clientQuery(gh *github.Client, query string) error {

	// Set the github search page we are currently searching.
	page := 1

	// Set the maximum page of github results for a given query.
	maxPage := math.MaxInt32

	// Set the github search options for the query.
	opts := github.SearchOptions{
		Sort:  "stars",
		Order: "desc",
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	}

	// Loop over pages in the github results so we can gather all
	// of the repositories in the search results.
	for page <= maxPage {

		// We can utilized the Search.Repositories method to Query github
		// for a particular page of results.
		opts.Page = page
		result, response, err := gh.Search.Repositories(query, &opts)
		if err != nil {
			return errors.Wrap(err, "Could not search Github result pages")
		}

		// Wait for the results.
		Wait(response)
		maxPage = response.LastPage

		// Then we loop over the repositories in the search results.
		for _, repo := range result.Repositories {

			// Extract the data of interest.
			name := *repo.FullName
			updatedAt := repo.UpdatedAt.String()
			createdAt := repo.CreatedAt.String()
			forks := *repo.ForksCount
			issues := *repo.OpenIssuesCount
			stars := *repo.StargazersCount
			size := *repo.Size

			// Print out the results for now.  However, this can be
			// redirected to a CSV file if desired.
			fmt.Printf("%s,%s,%s,%d,%d,%d,%d\n",
				name, updatedAt, createdAt, forks, issues, stars, size)
		}

		// Sleep for 10 seconds to stay within Github's rate limiting
		// constraints, then it increments the page and query again.
		time.Sleep(time.Second * 10)
		page++
	}

	return nil
}

// Wait waits to make sure we return the full github response.
func Wait(response *github.Response) {
	if response != nil && response.Remaining <= 1 {
		gap := time.Duration(response.Reset.Local().Unix() - time.Now().Unix())
		sleep := gap * time.Second
		if sleep < 0 {
			sleep = -sleep
		}
		time.Sleep(sleep)
	}
}
