package controllers

import (
	"log"
	"time"

	"localslackhook/views"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

type AppHomeController struct {
	EventHandler *socketmode.SocketmodeHandler
}

func NewAppHomeController(eventhandler *socketmode.SocketmodeHandler) AppHomeController {
	c := AppHomeController{
		EventHandler: eventhandler,
	}

	c.EventHandler.HandleEvents(
		slackevents.AppHomeOpened,
		c.publishHomeTabView,
	)

	c.EventHandler.HandleInteractionBlockAction(
		views.AddStickieNoteActionID,
		c.openCreateStickieNoteModal,
	)

	c.EventHandler.HandleInteraction(
		slack.InteractionTypeViewSubmission,
		c.createStickieNote,
	)

	return c
}

func (c *AppHomeController) publishHomeTabView(evt *socketmode.Event, clt *socketmode.Client) {
	evt_api, _ := evt.Data.(slackevents.EventsAPIEvent)
	evt_app_home_opened, _ := evt_api.InnerEvent.Data.(slackevents.AppHomeOpenedEvent)

	view := views.AppHomeTabView()

	_, err := clt.PublishView(evt_app_home_opened.User, view, "")

	if err != nil {
		log.Printf("Error publishHomeTabView: %v", err)
	}
}

func (c *AppHomeController) openCreateStickieNoteModal(evt *socketmode.Event, clt *socketmode.Client) {
	interaction := evt.Data.(slack.InteractionCallback)

	clt.Ack(*evt.Request)

	view := views.CreateStickieNoteModal()

	_, err := clt.OpenView(interaction.TriggerID, view)

	if err != nil {
		log.Printf("Error openCreateStickieNoteModal: %v", err)
	}
}

func (c *AppHomeController) createStickieNote(evt *socketmode.Event, clt *socketmode.Client) {
	view_submission := evt.Data.(slack.InteractionCallback)

	clt.Ack(*evt.Request)

	note := views.StickieNote{
		Description: view_submission.View.State.Values[views.ModalDescriptionBlockID][views.ModalDescriptionActionID].Value,
		Color:       view_submission.View.State.Values[views.ModalColorBlockID][views.ModalColorActionID].SelectedOption.Value,
		Timestamp:   time.Unix(time.Now().Unix(), 0).String(),
	}

	view := views.AppHomeCreateStickieNote(note)

	_, err := clt.PublishView(view_submission.User.ID, view, "")

	if err != nil {
		log.Printf("Error createStickieNote: %v", err)
	}
}
