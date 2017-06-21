package main

import (
	slackbot "github.com/BeepBoopHQ/go-slackbot"
	"github.com/DevinCarr/goarxiv"
	"github.com/nlopes/slack"
	"golang.org/x/net/context"
	"os"
	"strings"
)

var catmap = map[string]string{
	"AR": "Architecture",
	"AI": "Artificial Intelligence",
	"CL": "Computation and Language",
	"CC": "Computational Complexity",
	"CE": "Computational Engineering; Finance; and Science",
	"CG": "Computational Geometry",
	"GT": "Computer Science and Game Theory",
	"CV": "Computer Vision and Pattern Recognition",
	"CY": "Computers and Society",
	"CR": "Cryptography and Security",
	"DS": "Data Structures and Algorithms",
	"DB": "Databases",
	"DL": "Digital Libraries",
	"DM": "Discrete Mathematics",
	"DC": "Distributed; Parallel; and Cluster Computing",
	"GL": "General Literature",
	"GR": "Graphics",
	"HC": "Human-Computer Interaction",
	"IR": "Information Retrieval",
	"IT": "Information Theory",
	"LG": "Learning",
	"LO": "Logic in Computer Science",
	"MS": "Mathematical Software",
	"MA": "Multiagent Systems",
	"MM": "Multimedia",
	"NI": "Networking and Internet Architecture",
	"NE": "Neural and Evolutionary Computing",
	"NA": "Numerical Analysis",
	"OS": "Operating Systems",
	"OH": "Other",
	"PF": "Performance",
	"PL": "Programming Languages",
	"RO": "Robotics",
	"SE": "Software Engineering",
	"SD": "Sound",
	"SC": "Symbolic Computation",
}

func main() {
	bot := slackbot.New(os.Getenv("SLACK_TOKEN"))

	toMe := bot.Messages(slackbot.DirectMessage, slackbot.DirectMention, slackbot.Mention).Subrouter()
	toMe.Hear("(?i)(hi|hello).*").MessageHandler(HelloHandler)
	bot.Hear("(?i)categories(.*)").MessageHandler(CategoriesHandler)
	bot.Hear("(?i)papers").MessageHandler(PapersHandler)
	bot.Run()
}

func HelloHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	bot.Reply(evt, "Oh hello!", slackbot.WithTyping)
}

func CategoriesHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	parts := strings.Fields(evt.Text)
	if len(parts) == 2 && parts[0] == "categories" && parts[1] != "help" {
		_, ok := catmap[parts[1]]
		if ok {
			s := goarxiv.New()
			s.AddQuery("search_query", "cat:cs."+parts[1])
			s.AddQuery("sortBy", "submittedDate")
			s.AddQuery("sortOrder", "descending")
			result, err := s.Get()
			if err != nil {
				bot.Reply(evt, "Hey, something broke. Try again?", slackbot.WithTyping)
			}
			for i := 0; i < 10; i++ {
				attachment := slack.Attachment{
					Title:      result.Entry[i].Title,
					AuthorName: result.Entry[i].Author.Name,
					Text:       result.Entry[i].Summary.Body,
					TitleLink:  result.Entry[i].Link[1].Href,
					Fallback:   result.Entry[i].Summary.Body,
					Color:      "#371dba",
				}

				attachments := []slack.Attachment{attachment}

				bot.ReplyWithAttachments(evt, attachments, slackbot.WithTyping)
			}
		} else {
			bot.Reply(evt, "Invalid category! Type \"categories help\" for instructions.", slackbot.WithTyping)
		}
	}
	if len(parts) == 2 && parts[0] == "categories" && parts[1] == "help" {
		bot.Reply(evt, "Looking for help?", slackbot.WithTyping)
		bot.Reply(evt, "The allowed categories are: ", slackbot.WithTyping)
		bot.Reply(evt, "AR (Architecture)\n AI (Artificial Intelligence)\n CL (Computation and Language)\n CC (Computational Complexity)\n CE (Computational Engineering; Finance; and Science)\n CG (Computational Geometry)\n GT (Computer Science and Game Theory)\n CV (Computer Vision and Pattern Recognition)\n CY (Computers and Society)\n CR (Cryptography and Security)\n DS (Data Structures and Algorithms)\n DB (Databases)\n DL (Digital Libraries)\n DM (Discrete Mathematics\n DC (Distributed; Parallel; and Cluster Computing)\n GL (General Literature)\n GR (Graphics)\n HC (Human-Computer Interaction)\n IR (Information Retrieval)\n IT (Information Theory)\n LG (Learning)\n LO (Logic in Computer Science)\n MS (Mathematical Software)\n MA (Multiagent Systems)\n MM (Multimedia)\n NI (Networking and Internet Architecture)\n NE (Neural and Evolutionary Computing)\n NA (Numerical Analysis)\n OS (Operating Systems)\n OH (Other)\n PF (Performance)\n PL (Programming Languages)\n RO (Robotics)\n SE (Software Engineering)\n SD (Sound)\n SC (Symbolic Computation)", slackbot.WithoutTyping)
	}
}

func PapersHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	s := goarxiv.New()
	s.AddQuery("search_query", "cat:cs.CV")
	s.AddQuery("sortBy", "submittedDate")
	s.AddQuery("sortOrder", "descending")
	result, err := s.Get()
	if err != nil {
		bot.Reply(evt, "Hey, something broke. Try again?", slackbot.WithTyping)
	}
	for i := 0; i < 10; i++ {
		attachment := slack.Attachment{
			Title:      result.Entry[i].Title,
			AuthorName: result.Entry[i].Author.Name,
			Text:       result.Entry[i].Summary.Body,
			TitleLink:  result.Entry[i].Link[1].Href,
			Fallback:   result.Entry[i].Summary.Body,
			Color:      "#371dba",
		}

		attachments := []slack.Attachment{attachment}

		bot.ReplyWithAttachments(evt, attachments, slackbot.WithTyping)
		//bot.Reply(evt, result.Entry[i].Link[1], slackbot.WithTyping)
	}
}
