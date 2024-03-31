package main

import (
	"fmt"
	"github.com/Hazegard/McSoLogs/config"
	logReader "github.com/Hazegard/McSoLogs/logReader"
	"github.com/Hazegard/McSoLogs/notifier"
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
	flag.StringVarP(&logFile, "log-file", "f", "", "Log File to monitor")
	flag.StringVarP(&configFile, "config", "c", "", "Config file containing the webhook url")
	flag.BoolVar(&debug, "debug", false, "Re-read the current file (used for debugging purpose)")
	flag.Parse()
	c, err := config.NewConfig(logFile, configFile, debug)
	if err != nil {
		fmt.Printf("error parsing configuration: %s", err)
		os.Exit(1)
	}
	return c
}
