// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/adonovan/gopl.io/ch4/github"
)

//!+
func main() {
	const monthDuration = 2.628e+15
	const yearDuration = 3.154e+16
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	var lastMonthIssues, lastYearIssues, olderIssues *github.IssuesSearchResult = new(github.IssuesSearchResult),
		new(github.IssuesSearchResult),
		new(github.IssuesSearchResult)
	for _, item := range result.Items {
		if time.Now().Sub(item.CreatedAt) < monthDuration {
			lastMonthIssues.Items = append(lastMonthIssues.Items, item)
		} else if time.Now().Sub(item.CreatedAt) < yearDuration {
			lastYearIssues.Items = append(lastYearIssues.Items, item)
		} else {
			olderIssues.Items = append(olderIssues.Items, item)
		}
		lastMonthIssues.TotalCount = len(lastMonthIssues.Items)
		lastYearIssues.TotalCount = len(lastYearIssues.Items)
		olderIssues.TotalCount = len(olderIssues.Items)
	}
	fmt.Println("Issues open less than a month ago:")
	for _, item := range lastMonthIssues.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println("Issues open less than a year ago:")
	for _, item := range lastYearIssues.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println("Older issues:")
	for _, item := range olderIssues.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
