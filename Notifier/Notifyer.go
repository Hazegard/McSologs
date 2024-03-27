package Notifier

import (
	"github.com/containrrr/shoutrrr"
	"mc_so_logs/config"
	"time"
)

type Notifier struct {
	Url string
}

func NewNotifier(c *config.Config) Notifier {
	return Notifier{Url: c.DiscordUrl}
}

func (n *Notifier) Notify(message string) error {
	err := shoutrrr.Send(n.Url, message)
	time.Sleep(500 * time.Millisecond)
	return err
}
