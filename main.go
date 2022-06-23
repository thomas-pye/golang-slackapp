package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"localslackhook/controllers"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func main() {

	err := godotenv.Load("./test_slack.env")

	appToken := os.Getenv("SLACK_APP_TOKEN")
	botToken := os.Getenv("SLACK_BOT_TOKEN")
	channelId := os.Getenv("SLACK_CHANNEL_ID")

	api := slack.New(
		botToken,
		slack.OptionDebug(true),
		slack.OptionAppLevelToken(appToken),
	)

	client := socketmode.New(
		api,
		socketmode.OptionDebug(true),
		socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
	)

	attachment := slack.Attachment{
		Pretext: "Bot Message",
		Text:    "Bot is active",
		Color:   "#36a64f",

		Fields: []slack.AttachmentField{
			{
				Title: "Date",
				Value: time.Now().String(),
			},
		},
	}

	_, timestamp, err := client.PostMessage(
		channelId,
		slack.MsgOptionAttachments(attachment),
	)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Message sent at %s", timestamp)

	socketmodeHandler := socketmode.NewSocketmodeHandler(client)

	controllers.NewAppHomeController(socketmodeHandler)

	socketmodeHandler.RunEventLoop()
}
