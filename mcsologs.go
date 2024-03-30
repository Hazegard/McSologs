package main

import (
	"fmt"
	"github.com/Hazegard/McSoLogs/config"
	logReader "github.com/Hazegard/McSoLogs/logReader"
	"github.com/Hazegard/McSoLogs/notifier"
	flag "github.com/spf13/pflag"
)

func main() {
	logFile := ""
	notifyUrlFIle := ""
	debug := false
	flag.StringVarP(&logFile, "log-file", "f", "", "Log File to monitor")
	flag.StringVarP(&notifyUrlFIle, "config", "c", "", "Config file containing the webhook url")
	flag.BoolVar(&debug, "reread", false, "Re-read the current file (used for debugging purpose)")
	flag.Parse()
	c, err := config.NewConfig(logFile, notifyUrlFIle, debug)
	if err != nil {
		panic(err)
	}
	discordNotifier := notifier.NewNotifier(c)
	mcLogReader := logReader.NewLogReader(c)

	messages := mcLogReader.Message()

	go mcLogReader.TailFile()

	for m := range messages {
		if debug {
			fmt.Println(m)
		} else {
			err := discordNotifier.Notify(m)
			if err != nil {
				fmt.Println(err)
			}

		}
	}
}
