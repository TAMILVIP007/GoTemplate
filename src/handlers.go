package src

import (
	"github.com/amarnathcjd/gogram/telegram"
)

func PmStart(m *telegram.NewMessage) error {
	if _, err := m.Reply("Hello there! I'm a bot made by @mybotsrealm"); err != nil {
		return err
	}
	return nil
}
