package main

import (
	"fmt"
	"github.com/Hazegard/mcsologs/config"
	"github.com/Hazegard/mcsologs/logReader"
	"github.com/Hazegard/mcsologs/notifier"
	flag "github.com/spf13/pflag"
	"os"
)

func main() {
	c := parseConfig()
	err, discordNotifier := notifier.NewNotifier(c)
	if err != nil {
		panic(err)
	}
	mcLogReader := logReader.NewLogReader(c)

	messages := mcLogReader.Message()

	go mcLogReader.TailFile()

	for m := range messages {
		if c.Debug {
			fmt.Println(m)
		} else {
			discordNotifier.Notify(m)
		}
	}
}

func parseConfig() *config.Config {
	logFile := ""
	configFile := ""
	debug := false
	flag.StringVarP(&logFile, "log-file", "f", "latest.log", "Log File to monitor (required)")
	flag.StringVarP(&configFile, "config", "c", "config.json", "Config file containing the webhook url (required)")
	flag.BoolVar(&debug, "debug", false, "Re-read the current file (used for debugging purpose)")
	flag.Parse()
	if logFile == "" {
		fmt.Println("error, log file required")
		os.Exit(1)
	}
	if configFile == "" {
		fmt.Println("error, config file required")
		os.Exit(1)
	}

	c, err := config.NewConfig(logFile, configFile, debug)
	if err != nil {
		fmt.Printf("error parsing configuration: %s", err)
		os.Exit(1)
	}
	return c
}
