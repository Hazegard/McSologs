package notifier

import (
	"errors"
	"fmt"
	"github.com/Hazegard/mcsologs/config"
	"github.com/Hazegard/mcsologs/structs/message"
	"github.com/containrrr/shoutrrr"
	"github.com/containrrr/shoutrrr/pkg/router"
	"github.com/containrrr/shoutrrr/pkg/types"
	"time"
)

// Notifier holds the shoutrrr service used to send discord notification
type Notifier struct {
	Urls    []string
	senders []*router.ServiceRouter
	debug   bool
}

// NewNotifier returns the Notifier struct used to send discord notification
// it returns an error if the discord notification URL cannot be parsed by the underlying shoutrrr
func NewNotifier(c *config.Config) (error, *Notifier) {
	var senders []*router.ServiceRouter
	var errs []error
	for _, url := range c.DiscordUrls {
		sender, err := shoutrrr.CreateSender(url)
		if err != nil {
			errs = append(errs, fmt.Errorf("error while parsing discord url (%s): %s", c.DiscordUrls, err))
			continue
		}
		senders = append(senders, sender)
	}
	if len(senders) == 0 {
		return fmt.Errorf("error, no valid webhook found"), nil
	}

	return errors.Join(errs...), &Notifier{
		Urls:    c.DiscordUrls,
		senders: senders,
		debug:   c.Debug,
	}

}

// Notify sends a discord notification to the url configured in the struct
// The notification is customized depending on the Message struct passed to
// It sleeps for 0.5 second as a simple mechanism to prevent being rate limited by discord (429 too many requests)
// TODO: this system could be improved to not always sleeps, but only sleep when the rate limit is about to happen
func (n *Notifier) Notify(message message.Message) {
	params := make(types.Params)
	params["Color"] = message.GetWHColor()
	params["title"] = message.GetTitle()
	if n.debug {
		fmt.Printf("%s (%+v)", message.GetMessage(), params)
		return
	}
	for _, sender := range n.senders {
		sender.Send(message.GetMessage(), &params)
	}
	time.Sleep(500 * time.Millisecond)
}
