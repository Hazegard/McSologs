package main

import (
	"github.com/Hazegard/McSoLogs/regexParser"
	"github.com/Hazegard/McSoLogs/structs/message"
	"github.com/nxadm/tail"
	"io"
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

func HandleLine(line string) message.Message {
	var mes message.Message
	mes = regexParser.ParseDeath(line)
	if mes != nil {
		return mes
	}

	mes = regexParser.ParseAdvancement(line)
	if !mes.IsEmpty() {
		return mes
	}

	mes = regexParser.ParseLogout(line)
	if !mes.IsEmpty() {
		return mes
	}

	mes = regexParser.ParseLogin(line)
	if !mes.IsEmpty() {
		return mes
	}
	return mes
}
