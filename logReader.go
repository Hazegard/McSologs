package main

import (
	"github.com/nxadm/tail"
	"io"
	"mc_so_logs/RegexParser"
	"mc_so_logs/structs/Message"
)

func Read(messageChan chan<- string) {
	t, err := tail.TailFile("./mc.logs", tail.Config{
		Follow:   true,
		ReOpen:   true,
		Location: &tail.SeekInfo{Offset: 0, Whence: io.SeekEnd},
	})
	if err != nil {
		panic(err)
	}

	// Print the text of each received line
	for line := range t.Lines {
		mes := HandleLine(line.Text)
		if mes.IsEmpty() != true {
			messageChan <- mes.GetMessage()
		}
	}
}

func HandleLine(line string) Message.Message {
	var mes Message.Message
	mes = RegexParser.ParseDeath(line)
	if mes != nil {
		return mes
	}

	mes = RegexParser.ParseAdvancement(line)
	if !mes.IsEmpty() {
		return mes
	}

	mes = RegexParser.ParseLogout(line)
	if !mes.IsEmpty() {
		return mes
	}

	mes = RegexParser.ParseLogin(line)
	if !mes.IsEmpty() {
		return mes
	}
	return mes
}
