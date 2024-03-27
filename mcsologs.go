package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"mc_so_logs/Notifier"
	"mc_so_logs/config"
)

func main() {
	logFile := ""
	notifyUrlFIle := ""
	flag.StringVarP(&logFile, "log-file", "f", "", "Log File to monitor")
	flag.StringVarP(&notifyUrlFIle, "config", "c", "", "Log File to monitor")
	flag.Parse()
	c, err := config.NewConfig(logFile, notifyUrlFIle)
	if err != nil {
		panic(err)
	}
	notifier := Notifier.NewNotifier(c)
	var messages chan string = make(chan string)
	defer close(messages)

	go Read(messages)

	for m := range messages {
		err := notifier.Notify(m)
		if err != nil {
			fmt.Println(err)
		}
	}
}
