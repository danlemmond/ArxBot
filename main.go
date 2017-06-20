package main

import (
	"os"
	slackbot "github.com/BeepBoopHQ/go-slackbot"
	"github.com/nlopes/slack"
	"golang.org/x/net/context"
	"github.com/DevinCarr/goarxiv"
)

func main() {
	bot := slackbot.New(os.Getenv("SLACK_TOKEN"))

	toMe := bot.Messages(slackbot.DirectMessage, slackbot.DirectMention, slackbot.Mention).Subrouter()
	toMe.Hear("(?i)(hi|hello).*").MessageHandler(HelloHandler)
	bot.Hear("(?i)how are you?(.*)").MessageHandler(HowAreYouHandler)
	bot.Hear("(?i)papers").MessageHandler(AttachmentsHandler)
	bot.Run()
}

func HelloHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	bot.Reply(evt, "Oh hello!", slackbot.WithTyping)
}

func HowAreYouHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	bot.Reply(evt, "A bit tired. Get it? A bit?", slackbot.WithTyping)
}

func AttachmentsHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	s := goarxiv.New()
	s.AddQuery("search_query", "cat:cs.CV")
	result, err := s.Get()
	if err != nil {
		bot.Reply(evt, "Hey, something broke. Try again?", slackbot.WithTyping)
	}
	for i := 0; i < 10; i++ {
		attachment := slack.Attachment {
			Title: 		result.Entry[i].Title,
			AuthorName:	result.Entry[i].Author.Name,
			Text:		result.Entry[i].Summary.Body,
			Fallback: 	result.Entry[i].Summary.Body,
			Color: 		"#371dba",
		}

	attachments := []slack.Attachment{attachment}

	bot.ReplyWithAttachments(evt, attachments, slackbot.WithTyping)
	//bot.Reply(evt, result.Entry[i].Link[1], slackbot.WithTyping)
	}
}