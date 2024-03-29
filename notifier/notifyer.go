package notifier

import (
	"github.com/Hazegard/McSoLogs/config"
	"github.com/containrrr/shoutrrr"
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
