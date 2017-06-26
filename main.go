package main

import (
	slackbot "github.com/BeepBoopHQ/go-slackbot"
	"github.com/DevinCarr/goarxiv"
	"github.com/nlopes/slack"
	"golang.org/x/net/context"
	"os"
	"strings"
	"time"
)

func main() {
	bot := slackbot.New(os.Getenv("SLACK_TOKEN"))

	toMe := bot.Messages(slackbot.DirectMessage, slackbot.DirectMention, slackbot.Mention).Subrouter()
	go toMe.Hear("(?i)(hi|hello).*").MessageHandler(HelpHandler)
	go bot.Hear("(?i)author").MessageHandler(AuthorHandler)
	go bot.Hear("(?i)categories(.*)").MessageHandler(CategoriesHandler)
	go bot.Hear("(?i)arxbot(.*)").MessageHandler(HelpHandler)
	go bot.Hear("(?i)title(.*)").MessageHandler(TitleHandler)
	bot.Run()
}

//HelpHandler returns results and options for Arxbot, including available commands
func HelpHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	parts := strings.Fields(evt.Text)
	if len(parts) == 2 && parts[0] == "arxbot" && parts[1] == "help" {
		bot.Reply(evt, "Hey, thanks for using Arxbot, a paper retrieval bot for Slack", slackbot.WithTyping)
		bot.Reply(evt, "Arxbot is a bot for Arxiv that allows user to input their own search parameters and receive results.", slackbot.WithTyping)
		bot.Reply(evt, "The current available commands are author, title, and categories.", slackbot.WithTyping)
		bot.Reply(evt, "Type '[command] help' to get more information about a command, ex. author help", slackbot.WithTyping)
	}
	if len(parts) == 1 && parts[0] == "arxbot" {
		bot.Reply(evt, "Please use [arxbot help] for assistance using Arxbot.", slackbot.WithTyping)
	}
}

//TitleHandler allows users to query articles by title
func TitleHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	parts := strings.Fields(evt.Text)
	if len(parts) >= 2 && parts[0] == "title" && parts[1] != "help" {
		strjn := strings.Join(parts[1:], "%20")
		queryparam := "ti:\"" + strjn + "\""
		go QueryBuilder(ctx, bot, evt, queryparam)
		if len(parts) == 2 && parts[0] == "title" && parts[1] == "help" {
			bot.Reply(evt, "The Title query allows users to query Arxiv by article title", slackbot.WithTyping)
			bot.Reply(evt, "The command is used by typing:\ntitle [title of article]", slackbot.WithTyping)
		}
	}
}

//CategoriesHandler returns a list of the most recent 5 papers given a category and subcategory.
func CategoriesHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	var parts2var string
	parts := strings.Fields(evt.Text)
	if len(parts) == 3 {
		parts2var = parts[2]
	}
	if len(parts) == 3 && parts[0] == "categories" && parts[1] != "help" {
		_, ok := Catmap[parts[1]]
		if ok {
			queryparam := "cat:" + parts[1] + "." + parts[2]
			go QueryBuilder(ctx, bot, evt, queryparam)
		} else {
			bot.Reply(evt, "Sorry, invalid category or subcategory!", slackbot.WithTyping)
		}
	}
	if len(parts) == 2 && parts[0] == "categories" && parts[1] != "help" {
		_, ok := Primmap[parts[1]]
		if ok {
			queryparam := "cat:" + parts[1]
			go QueryBuilder(ctx, bot, evt, queryparam)
		} else {
			bot.Reply(evt, "Your query failed. Please verify that the information you entered is accurate.", slackbot.WithTyping)
			bot.Reply(evt, "Please be aware that Astrophysics, General Relativity and Quantum Cosmology, the High Energy Physics family, Mathematical Physics, Nuclear Experiment/Nuclear Theory, and Quantum Theory do NOT have subcategories.", slackbot.WithTyping)

		}
	}
	if len(parts) == 2 && parts[0] == "categories" && parts[1] == "help" {
		bot.Reply(evt, "The categories function will allow you to parse for papers by category, such as math, physics, or computer science.", slackbot.WithTyping)
		bot.Reply(evt, "Query format is: categories [primary] [secondary]", slackbot.WithTyping)
		bot.Reply(evt, "For example, [categories math LO] will return the 5 most recent Logic papers published to Arxiv.", slackbot.WithTyping)
		bot.Reply(evt, "To see all primary categores which DO NOT have secondary categories, type: categories help soloprimary.", slackbot.WithTyping)
		bot.Reply(evt, "To see all primary categories which DO have secondary categories, type: categories help primary.", slackbot.WithTyping)
		bot.Reply(evt, "To see all available secondary categories for a given primary category, type: categories help [primary], ex: categories help nlin.", slackbot.WithTyping)
	}
	switch parts2var {
	case "math":
		for k, v := range Mathmap {
			bot.Reply(evt, "Secondary Category: "+"\""+k+"\""+" Topic: "+v, slackbot.WithTyping)
			time.Sleep(0)
		}
	case "nlin":
		for k, v := range Nlinmap {
			bot.Reply(evt, "Secondary Category: "+"\""+k+"\""+" Topic: "+v, slackbot.WithTyping)
			time.Sleep(0)
		}
	case "q-bio":
		for k, v := range Qbiomap {
			bot.Reply(evt, "Secondary Category: "+"\""+k+"\""+" Topic: "+v, slackbot.WithTyping)
			time.Sleep(0)
		}
	case "stat":
		for k, v := range Statmap {
			bot.Reply(evt, "Secondary Category: "+"\""+k+"\""+" Topic: "+v, slackbot.WithTyping)
			time.Sleep(0)
		}
	case "cs":
		for k, v := range CSmap {
			bot.Reply(evt, "Secondary Category: "+"\""+k+"\""+" Topic: "+v, slackbot.WithTyping)
			time.Sleep(0)
		}
	case "cond-mat":
		for k, v := range Condmap {
			bot.Reply(evt, "Secondary Category: "+"\""+k+"\""+" Topic: "+v, slackbot.WithTyping)
			time.Sleep(0)
		}
	case "physics":
		for k, v := range Physmap {
			bot.Reply(evt, "Secondary Category: "+"\""+k+"\""+" Topic: "+v, slackbot.WithTyping)
			time.Sleep(0)
		}
	case "primary":
		for k, v := range Catmap {
			bot.Reply(evt, "Category: "+"\""+k+"\""+" Description: "+v, slackbot.WithTyping)
			time.Sleep(0)
		}
	case "soloprimary":
		for k, v := range Primmap {
			bot.Reply(evt, "Category: "+"\""+k+"\""+" Description: "+v, slackbot.WithTyping)
			time.Sleep(0)
		}
	default:
		break
	}
}

//AuthorHandler returns the papers written by a given author, submitted by the user.
func AuthorHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	parts := strings.Fields(evt.Text)
	if len(parts) == 2 && parts[0] == "author" && parts[1] != "help" {
		queryparam := "au:" + parts[1]
		go QueryBuilder(ctx, bot, evt, queryparam)
	}
	if len(parts) == 3 && parts[0] == "author" {
		a := []rune(parts[1])
		queryparam := "au:" + parts[2] + "_" + string(a[0])
		QueryBuilder(ctx, bot, evt, queryparam)
	}
	if len(parts) == 2 && parts[0] == "author" && parts[1] == "help" {
		bot.Reply(evt, "The author command allows you to search for authors by last name, or first and last name.", slackbot.WithTyping)
		bot.Reply(evt, "The two uses are: author [lastname] or author [first] [last]", slackbot.WithTyping)
	}
}

//QueryBuilder builds and returns an Arxiv query.
func QueryBuilder(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent, s string) {
	query := goarxiv.New()
	query.AddQuery("search_query", s)
	query.AddQuery("sortBy", "submittedDate")
	query.AddQuery("sortOrder", "descending")
	query.AddQuery("max_results", "5")
	result, err := query.Get()
	if err != nil {
		bot.Reply(evt, "Sorry, there was an error. Try again!", slackbot.WithTyping)
	}
	if len(result.Entry) == 0 {
		bot.Reply(evt, "Your query returned 0 results! Please be sure that your query information is correct!", slackbot.WithTyping)
	}
	for i := 0; i < len(result.Entry); i++ {
		strtp := string(result.Entry[i].Published)
		attachment := slack.Attachment{
			Title:      result.Entry[i].Title,
			AuthorName: result.Entry[i].Author.Name,
			Text:       result.Entry[i].Summary.Body,
			TitleLink:  result.Entry[i].Link[1].Href,
			Fallback:   result.Entry[i].Summary.Body,
			Footer:     "Published " + strtp,
			Color:      "#371dba",
		}

		attachments := []slack.Attachment{attachment}
		bot.ReplyWithAttachments(evt, attachments, slackbot.WithTyping)
	}
}
