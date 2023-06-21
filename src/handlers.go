package src

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func PmStart(b *gotgbot.Bot, ctx *ext.Context) error {
	if ctx.EffectiveChat.Type == "private" {
		ctx.EffectiveMessage.Reply(b, "Hello there! I'm a sticker bot made by @mybotsrealm.", nil)
	}
	return nil
}
