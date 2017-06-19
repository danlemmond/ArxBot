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
			
		}
	}
}