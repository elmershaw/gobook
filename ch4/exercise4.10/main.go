package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	orderByTime := make([][]github.Issue, 3)
	for _, item := range result.Items {
		if item.CreatedAt.After(time.Now().AddDate(0, -1, 0)) {
			orderByTime[0] = append(orderByTime[0], *item)
		} else if item.CreatedAt.After(time.Now().AddDate(-1, 0, 0)) {
			orderByTime[1] = append(orderByTime[1], *item)
		} else {
			orderByTime[2] = append(orderByTime[2], *item)
		}
	}
	for _, i := range orderByTime {
		for _, item := range i {
			fmt.Printf("#%-5d %9.9s %.55s %s\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
	}
}
