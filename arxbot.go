package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
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
			if len(parts) == 3 && parts[1] == "stock" {
				go func(m Message) {
					m.Text = getQuote(parts[2])
					postMessage(ws, m)
				}(m)

			} else {
				m.text = fmt.Sprintf("sorry, that does not compute\n")
				postMessage(ws, m)
			}
		}
	}
}

func getQuote(sym string) string {
	sym = strings.ToUpper(sym)
	url := fmt.Sprintf("https://download.finance.yahoo.com/d/quotes.csv?s=%s&f=nsl1op&e=.csv", sym)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	rows, err := csv.NewReader(resp.Body).ReadAll()
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	if len(rows) >= 1 && len(rows[0]) == 5 {
		return fmt.Sprintf("%s (%s) is trading at $%s" rows[0][0], rows [0][1], rows[0][2])
	}
	return fmt.Sprintf("Unknown response format (symbol was \"%s\")", sym)
}