package main

import (
	"GotgTemplate/src"

	"github.com/amarnathcjd/gogram/telegram"
)

func main() {
	client, _ := telegram.NewClient(telegram.ClientConfig{
		AppID:    src.Envars.AppId,
		AppHash:  src.Envars.AppHash,
		LogLevel: telegram.LogInfo,
		Session: "./bot.session",
	})
	if err := client.Connect(); err != nil {
		panic(err)
	}
	if err := client.LoginBot(src.Envars.Token); err != nil {
		panic(err)
	}
	client.AddMessageHandler("/start", src.PmStart)
	client.Idle()
}
