package main

import (
	"fmt"
	slackbot "github.com/BeepBoopHQ/go-slackbot"
	"github.com/DevinCarr/goarxiv"
	"github.com/nlopes/slack"
	"golang.org/x/net/context"
	"os"
	"strings"
)

//These Categories have subcategories, such as Atom Physics or Machine Learning.
//They will have a separate map in order to facilitate improved error checking.
var catmap = map[string]string{
	"stat":     "Statistics",
	"q-bio":    "Quantitative Biology",
	"cs":       "Computer Science",
	"nlin":     "Nonlinear Sciences",
	"math":     "Math",
	"cond-mat": "Physics - Mat",
	"physics":  "Physics",
}

//These categories have NO SECONDARY CATEGORIES. This means that any value passed after the input will crash Arxbot
//We'll do error checking against his map when doing our Primary Category check.

var primmap = map[string]string{
	"astro-ph": "Astrophysics",
	"gr-qc":    "General Relativity",
	"hep-ex":   "High Energy Physics - Experiment",
	"hep-lat":  "High Energy Physics - Lattice",
	"hep-ph":   "High Energy Physics - Phenomenology",
	"hep-th":   "High Energy Physics - Theory",
	"math-ph":  "Mathematical Physics",
	"nucl-ex":  "Nuclear Experiment",
	"nucl-th":  "Nuclear Theory",
	"quant-ph": "Quantum Physics",
}

//These following suffix maps exist because I thought I would use them for error checking during the categories call.
//Luckily I didn't have to, but I'm loathe to delete them.
var statmap = map[string]string{
	"AP": "Applications",
	"CO": "Computation",
	"ML": "Machine Learning",
	"ME": "Methodology",
	"TH": "Theory",
}

var qbiomap = map[string]string{
	"BM": "Biomolecules",
	"CB": "Cell Behavior",
	"GN": "Genomics",
	"MN": "Molecular Networks",
	"NC": "Neurons and Cognition",
	"OT": "Other",
	"PE": "Populations and Evolution",
	"QM": "Quantitative Methods",
	"SC": "Subcellular Processes",
	"TO": "Tissues and Organs",
}

var nlinmap = map[string]string{
	"AO": "Adaptation and Self-Organizing Systems",
	"CG": "Cellular Atuomata and Lattice Gasses",
	"CD": "Chaotic Dynamics",
	"SI": "Exactly Solvable and Integrable Systems",
	"PS": "Pattern Formation and Solitons",
}

var mathmap = map[string]string{
	"AG": "Algebraic Geometry",
	"AT": "Algebraic Topology",
	"AP": "Analysis of PDEs",
	"CT": "Category Theory",
	"CA": "Classical Analysis and ODEs",
	"CO": "Combinatorics",
	"AC": "Commutative Algebra",
	"CV": "Complex Variables",
	"DG": "Differential Geometry",
	"DS": "Dynamical Systems",
	"FA": "Functional Analysis",
	"GM": "General Mathematics",
	"GN": "General Topology",
	"GT": "Geometric Topology",
	"GR": "Group Theory",
	"HO": "History and Overview",
	"IT": "Information Theory",
	"KT": "K-Theory and Homology",
	"LO": "Logic",
	"MP": "Mathematical Physics",
	"MG": "Metric Geometry",
	"NT": "Number Theory",
	"NA": "Numerical Analysis",
	"OA": "Operator Algebras",
	"OC": "Optimization and Control",
	"PR": "Probability",
	"QA": "Quantum Algebra",
	"RT": "Representation Theory",
	"RA": "Rings and Algebras",
	"SP": "Spectral Theory",
	"ST": "Statistics",
	"SG": "Symplectic Geometry",
}

var condmap = map[string]string{
	"dis-nn":    "Disordered Systems and Neural Networks",
	"mes-hall":  "Mesoscopic Systems and Quantum Hall Effect",
	"mtrl-sci":  "Materials Science",
	"other":     "Other",
	"soft":      "Soft Condensed Matter",
	"stat-mech": "Statistical Mechanics",
	"str-el":    "Strongly Correlated Electrons",
	"supr-con":  "Superconductivity",
}

var physmap = map[string]string{
	"acc-ph":   "Accelerator Physics",
	"ao-ph":    "Atmospheric and Oceanic Physics",
	"atom-ph":  "Atomic Physics",
	"atm-clus": "Atomic and Molecular Clusters",
	"bio-ph":   "Biological Physics",
	"chem-ph":  "Chemical Physics",
	"class-ph": "Classical Physics",
	"comp-ph":  "Computational Physics",
	"data-an":  "Data Analysis; Statistics and Probability",
	"flu-dyn":  "Fluid Dynamics",
	"gen-ph":   "General Physics",
	"geo-ph":   "Geophysics",
	"hist-ph":  "History of Physics",
	"ins-det":  "Instrumentation and Detectors",
	"med-ph":   "Medical Physics",
	"optics":   "Optics",
	"ed-ph":    "Physics Education",
	"soc-ph":   "Physics and Society",
	"plasm-ph": "Plasma Physics",
	"pop-ph":   "Popular Physics",
	"space-ph": "Space Physics",
}

var csmap = map[string]string{
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
	go toMe.Hear("(?i)(hi|hello).*").MessageHandler(HelpHandler)
	go bot.Hear("(?i)cs(.*)").MessageHandler(CSCategoriesHandler)
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
		bot.Reply(evt, "Hey, thanks for using Arxbot, an Arxiv parser for Slack!", slackbot.WithTyping)
		bot.Reply(evt, "Arxbot is a dynamic parser for Arxiv that allows user to input their own search parameters and receive results.", slackbot.WithTyping)
		bot.Reply(evt, "The current available commands are author, cs, title, and categories.", slackbot.WithTyping)
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
		s := goarxiv.New()
		s.AddQuery("search_query", "ti:\""+strjn+"\"")
		s.AddQuery("sortBy", "submittedDate")
		s.AddQuery("sortOrder", "descending")
		s.AddQuery("max_results", "5")
		fmt.Println(s.Query)
		result, err := s.Get()
		if err != nil {
			bot.Reply(evt, "Something broke! Please try again.", slackbot.WithTyping)
		}
		for i := 0; i < len(result.Entry); i++ {
			strtm := string(result.Entry[i].Published)
			attachment := slack.Attachment{
				Title:      result.Entry[i].Title,
				AuthorName: result.Entry[i].Author.Name,
				Text:       result.Entry[i].Summary.Body,
				TitleLink:  result.Entry[i].Link[1].Href,
				Fallback:   result.Entry[i].Summary.Body,
				Footer:     strtm,
				Color:      "#371dba",
			}

			attachments := []slack.Attachment{attachment}

			bot.ReplyWithAttachments(evt, attachments, slackbot.WithTyping)
		}
	}
	if len(parts) == 2 && parts[0] == "title" && parts[1] == "help" {
		bot.Reply(evt, "The Title query allows users to query Arxiv by article title", slackbot.WithTyping)
		bot.Reply(evt, "The command is used by typing:\ntitle [title of article]", slackbot.WithTyping)
	}
}

//CSCategoriesHandler returns a list of the 5 most recent papers in a CS category. Help returns a list of options.
func CSCategoriesHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	parts := strings.Fields(evt.Text)
	if len(parts) == 2 && parts[0] == "cs" && parts[1] != "help" {
		_, ok := csmap[parts[1]]
		if ok {
			s := goarxiv.New()
			s.AddQuery("search_query", "cat:cs."+parts[1])
			s.AddQuery("sortBy", "submittedDate")
			s.AddQuery("sortOrder", "descending")
			s.AddQuery("max_results", "5")
			result, err := s.Get()
			if err != nil {
				bot.Reply(evt, "Hey, something broke. Try again?", slackbot.WithTyping)
			}
			if len(result.Entry) == 0 {
				bot.Reply(evt, "Your query returned 0 results! Please be sure your query information is correct.", slackbot.WithTyping)
			}
			for i := 0; i < len(result.Entry); i++ {
				strtm := string(result.Entry[i].Published)
				attachment := slack.Attachment{
					Title:      result.Entry[i].Title,
					AuthorName: result.Entry[i].Author.Name,
					Text:       result.Entry[i].Summary.Body,
					TitleLink:  result.Entry[i].Link[1].Href,
					Fallback:   result.Entry[i].Summary.Body,
					Footer:     strtm,
					Color:      "#371dba",
				}

				attachments := []slack.Attachment{attachment}

				bot.ReplyWithAttachments(evt, attachments, slackbot.WithTyping)
			}
		} else {
			bot.Reply(evt, "Invalid category! Type \"cs help\" for instructions.", slackbot.WithTyping)
		}
	}
	if len(parts) == 2 && parts[0] == "cs" && parts[1] == "help" {
		bot.Reply(evt, "Looking for help?", slackbot.WithTyping)
		bot.Reply(evt, "The allowed categories are: ", slackbot.WithTyping)
		bot.Reply(evt, "AR (Architecture)\n AI (Artificial Intelligence)\n CL (Computation and Language)\n CC (Computational Complexity)\n CE (Computational Engineering; Finance; and Science)\n CG (Computational Geometry)\n GT (Computer Science and Game Theory)\n CV (Computer Vision and Pattern Recognition)\n CY (Computers and Society)\n CR (Cryptography and Security)\n DS (Data Structures and Algorithms)\n DB (Databases)\n DL (Digital Libraries)\n DM (Discrete Mathematics)\n DC (Distributed; Parallel; and Cluster Computing)\n GL (General Literature)\n GR (Graphics)\n HC (Human-Computer Interaction)\n IR (Information Retrieval)\n IT (Information Theory)\n LG (Learning)\n LO (Logic in Computer Science)\n MS (Mathematical Software)\n MA (Multiagent Systems)\n MM (Multimedia)\n NI (Networking and Internet Architecture)\n NE (Neural and Evolutionary Computing)\n NA (Numerical Analysis)\n OS (Operating Systems)\n OH (Other)\n PF (Performance)\n PL (Programming Languages)\n RO (Robotics)\n SE (Software Engineering)\n SD (Sound)\n SC (Symbolic Computation)", slackbot.WithoutTyping)
	}
}

//CategoriesHandler returns a list of the most recent 5 papers given a category and subcategory.
func CategoriesHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	parts := strings.Fields(evt.Text)
	if len(parts) == 3 && parts[0] == "categories" && parts[1] != "help" {
		_, ok := catmap[parts[1]]
		if ok {
			s := goarxiv.New()
			s.AddQuery("search_query", "cat:"+parts[1]+"."+parts[2])
			s.AddQuery("sortBy", "submittedDate")
			s.AddQuery("sortOrder", "descending")
			s.AddQuery("max_results", "5")
			result, err := s.Get()
			if err != nil {
				bot.Reply(evt, "Something went wrong. Please try again.", slackbot.WithTyping)
			}
			if len(result.Entry) == 0 {
				bot.Reply(evt, "Your query returned 0 results, which is most likely an error. Please be sure your subcategory is correct!", slackbot.WithTyping)
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
		} else {
			bot.Reply(evt, "Sorry, invalid category or subcategory!", slackbot.WithTyping)
		}
	}
	if len(parts) == 2 && parts[0] == "categories" && parts[1] != "help" {
		_, ok := primmap[parts[1]]
		if ok {
			s := goarxiv.New()
			s.AddQuery("search_query", "cat:"+parts[1])
			s.AddQuery("sortBy", "submittedDate")
			s.AddQuery("sortOrder", "descending")
			s.AddQuery("max_results", "5")
			result, err := s.Get()
			if err != nil {
				bot.Reply(evt, "There was an error. Please try again!", slackbot.WithTyping)
			}
			if len(result.Entry) == 0 {
				bot.Reply(evt, "Your query returned 0 results! Please make sure that your query information is correct.", slackbot.WithTyping)
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
		} else {
			bot.Reply(evt, "Your query failed. Please verify that the information you entered is accurate.", slackbot.WithTyping)
			bot.Reply(evt, "Please be aware that Astrophysics, General Relativity and Quantum Cosmology, the High Energy Physics family, Mathematical Physics, Nuclear Experiment/Nuclear Theory, and Quantum Theory do NOT have subcategories.", slackbot.WithTyping)

		}
	}
	if len(parts) == 2 && parts[0] == "categories" && parts[1] == "help" {
		bot.Reply(evt, "The categories function will allow you to parse for papers for categories outside of Computer Science, such as math or physics.", slackbot.WithTyping)
		bot.Reply(evt, "Query format is: categories [primary] [secondary]", slackbot.WithTyping)
		bot.Reply(evt, "For example, [categories math LO] will return the 5 most recent Logic papers published to Arxiv.", slackbot.WithTyping)
	}
}

//AuthorHandler returns the papers written by a given author, submitted by the user.
func AuthorHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	parts := strings.Fields(evt.Text)
	if len(parts) == 2 && parts[0] == "author" && parts[1] != "help" {
		s := goarxiv.New()
		s.AddQuery("search_query", "au:"+parts[1])
		s.AddQuery("sortBy", "submittedDate")
		s.AddQuery("sortOrder", "descending")
		s.AddQuery("max_results", "5")
		result, err := s.Get()
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
	if len(parts) == 3 && parts[0] == "author" {
		a := []rune(parts[1])
		s := goarxiv.New()
		s.AddQuery("search_query", "au:"+parts[2]+"_"+string(a[0]))
		s.AddQuery("sortBy", "submittedDate")
		s.AddQuery("sortOrder", "descending")
		s.AddQuery("max_results", "5")
		result, err := s.Get()
		if err != nil {
			bot.Reply(evt, "Sorry, there was an error. Try again!", slackbot.WithTyping)
		}
		if len(result.Entry) == 0 {
			bot.Reply(evt, "Your query returned 0 results! Please be sure your query information is correct.", slackbot.WithTyping)
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
	if len(parts) == 2 && parts[0] == "author" && parts[1] == "help" {
		bot.Reply(evt, "The author command allows you to search for authors by last name, or first and last name.", slackbot.WithTyping)
		bot.Reply(evt, "The two uses are: author [lastname] or author [first] [last]", slackbot.WithTyping)
	}
}
