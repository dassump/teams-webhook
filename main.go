package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	teams "github.com/atc0005/go-teams-notify/v2"
	"github.com/urfave/cli/v2"
)

var (
	app_name      string = "teams-webhook"
	app_version   string = "dev"
	app_usage     string = "Send a message to Microsoft Teams® via Webhook"
	app_copyright string = "GNU General Public License"
)

const (
	author_name  string = "Daniel Dias de Assumpção"
	author_email string = "dassump@gmail.com"

	flag_color_name     string = "color"
	flag_color_alias    string = "c"
	flag_color_usage    string = "card theme color"
	flag_color_default  string = "#DF813D"
	flag_color_required bool   = false

	flag_message_name     string = "message"
	flag_message_alias    string = "m"
	flag_message_usage    string = "card text or file path to get content"
	flag_message_required bool   = true

	flag_title_name     string = "title"
	flag_title_alias    string = "t"
	flag_title_usage    string = "card title"
	flag_title_required bool   = true

	flag_url_name     string = "url"
	flag_url_alias    string = "u"
	flag_url_usage    string = "teams webhook url"
	flag_url_required bool   = true
)

func main() {
	app := cli.NewApp()

	app.Name = app_name
	app.Usage = app_usage
	app.Version = app_version
	app.Authors = []*cli.Author{{Name: author_name, Email: author_email}}
	app.Copyright = app_copyright
	app.HideHelp = true

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     flag_color_name,
			Aliases:  []string{flag_color_alias},
			Usage:    flag_color_usage,
			Value:    flag_color_default,
			Required: flag_color_required,
		},
		&cli.StringFlag{
			Name:     flag_message_name,
			Aliases:  []string{flag_message_alias},
			Usage:    flag_message_usage,
			Required: flag_message_required,
		},
		&cli.StringFlag{
			Name:     flag_title_name,
			Aliases:  []string{flag_title_alias},
			Usage:    flag_title_usage,
			Required: flag_title_required,
		},
		&cli.StringFlag{
			Name:     flag_url_name,
			Aliases:  []string{flag_url_alias},
			Usage:    flag_url_usage,
			Required: flag_url_required,
		},
	}

	app.Action = func(c *cli.Context) error {
		card := teams.NewMessageCard()
		card.ThemeColor = c.String(flag_color_name)
		card.Title = c.String(flag_title_name)

		file, err := ioutil.ReadFile(c.String(flag_message_name))
		if err == nil {
			card.Text = string(file)
		} else {
			card.Text = c.String(flag_message_name)
		}

		return teams.NewClient().Send(c.String(flag_url_name), card)
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("\n%v\n", err)
		os.Exit(1)
	}
}
