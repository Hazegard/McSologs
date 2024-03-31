package notifier

import (
	"github.com/Hazegard/McSoLogs/config"
	"github.com/Hazegard/McSoLogs/structs/message"
	"github.com/containrrr/shoutrrr"
	"github.com/containrrr/shoutrrr/pkg/types"
	"time"
)

type Notifier struct {
	Url string
}

func NewNotifier(c *config.Config) Notifier {
	return Notifier{Url: c.DiscordUrl}
}

func (n *Notifier) Notify(message message.Message) error {
	sender, err := shoutrrr.CreateSender(n.Url)
	if err != nil {
		return err
	}
	params := make(types.Params)
	params["Color"] = message.GetWHColor()
	params["title"] = message.GetTitle()
	sender.Send(message.GetMessage(), &params)
	time.Sleep(500 * time.Millisecond)
	return nil
}
