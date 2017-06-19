package main

import (
	"fmt"
	"log"
	//"net/http"
	"os"
	"strings"
	"github.com/DevinCarr/goarxiv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: arxbot slack-bot-token")
		os.Exit(1)
	}

	ws, id := slackConnect(os.Args[1])
	fmt.Println("Arxbot running. ^C to quit.")

	for {
		m, err := getMessage(ws)
		if err != nil {
			log.Fatal(err)
		}

		if m.Type == "message" && strings.HasPreix(m.Text, "<@"+id+">") {
			parts := strings.Fields(m.Text)
			if len(parts) == 3 && parts[1] == "papers" {
				go func(m Message) {
					m.Text = getPapers(parts[2])
					postMessage(ws, m)
				}(m)

			} else {
				fmt.Println("Query failed. Please use @arxivbot help for assistance.")
			}
			if len(parts) == 5 && parts[1] == "author" {
				go func(m Message) {
					m.Text = authorSearch(parts[2:3], parts[4])
					postMessage(ws, m)
				}(m)
			} else {
				fmt.Println("Query failed. Please use @arxivbot help for assistance.")
			}
			if len(parts) == 4 && parts[1] == "topic" {
				go func(m Message) {
					m.Text = topicSearch(parts[2], parts[3])
					postMessage(ws, m)
				}(m)
			} else {
				fmt.Println("Query failed. Please use @arxivbot help for assistance.")
			}

		}
	}
}

func getPapers(n int) string {
	newsearch = goarxiv.New()
	newsearch.AddQuery("search_query", "cat:cs.CV")
	resp, err := newsearch.Get()
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	for i := 0; i < n; i++ {
		fmt.Println(resp.Entry[i].Title, resp.Entry[i].Summary, resp.Entry[i].Author, resp.Entry[i].Link)
	}
}

func authorSearch (s string, n int) string {
	newsearch = goarxiv.New()
	newsearch.AddQuery("search_query", "au: %s", s)
	resp, err := newsearch.Get()
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	for i := 0; i < n; i++ {
		fmt.Println(resp.Entry[i].Title, resp.Entry[i].Summary, resp.Entry[i].Author, resp.Entry[i].Link)
	}
}

func topicSearch (s string, n int) {
	newsearch := goarxiv.New()
	newsearch.AddQuery("search_query", "cat:%s", s)
	resp, err := newsearch.Get()
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	for i := 0; i < n; i++ {
		fmt.Println(resp.Entry[i].Title, resp.Entry[i].Summary, resp.Entry[i].Author, resp.Entry[i].Link)
	}
}

//func keywordSearch (s []string) {
//
//}
