package main

import (
	"GotgTemplate/src"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func main() {
	b, err := gotgbot.NewBot(src.Envars.Token, &gotgbot.BotOpts{
		Client: http.Client{},
		DefaultRequestOpts: &gotgbot.RequestOpts{
			Timeout: gotgbot.DefaultTimeout,
			APIURL:  gotgbot.DefaultAPIURL,
		},
	})
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}
	updater := ext.NewUpdater(&ext.UpdaterOpts{
		Dispatcher: ext.NewDispatcher(&ext.DispatcherOpts{
			Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
				log.Println("an error occurred while handling update:", err.Error())
				return ext.DispatcherActionNoop
			},
			MaxRoutines: ext.DefaultMaxRoutines,
		}),
	})
	dispatcher := updater.Dispatcher
	dispatcher.AddHandler(handlers.NewCommand("start", src.PmStart))
	if src.Envars.Webhook == "" {
		err = updater.StartPolling(b, &ext.PollingOpts{
			DropPendingUpdates: false,
			GetUpdatesOpts: gotgbot.GetUpdatesOpts{
				Timeout: 9,
				RequestOpts: &gotgbot.RequestOpts{
					Timeout: time.Second * 10,
				},
			},
		})
		if err != nil {
			panic("failed to start polling: " + err.Error())
		}
		log.Printf("%s has been started...\n Polling...\n", b.User.Username)
	} else {
		err = updater.StartWebhook(b, src.Envars.Token, ext.WebhookOpts{ListenAddr: fmt.Sprintf("localhost:%s", src.Envars.Port)})

		if err != nil {
			panic("failed to start webhook: " + err.Error())
		}
		err = updater.SetAllBotWebhooks(src.Envars.Webhook, &gotgbot.SetWebhookOpts{
			MaxConnections: 40,
			AllowedUpdates: []string{"message", "callback_query"},
		})
		if err != nil {
			panic("failed to set webhook: " + err.Error())
		}
		log.Printf("%s has been started...\n Webhook: %s\n", b.User.Username, src.Envars.Webhook)
	}
	updater.Idle()
}
