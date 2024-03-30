package logReader

import (
	"github.com/Hazegard/McSoLogs/config"
	"github.com/Hazegard/McSoLogs/regexParser"
	"github.com/Hazegard/McSoLogs/structs/message"
	"github.com/nxadm/tail"
	"io"
)

type LogReader struct {
	messageChan chan string
	doReadFile  bool
	tailFile    tail.Tail
	logFile     string
}

func NewLogReader(c *config.Config) *LogReader {
	return &LogReader{
		messageChan: make(chan string),
		doReadFile:  c.Debug,
		logFile:     c.LogFile,
	}
}

func (l *LogReader) TailFile() {
	t, err := tail.TailFile(l.logFile, l.generateTailConfig())
	if err != nil {
		panic(err)
	}
	for line := range t.Lines {
		mes := handleLine(line.Text)
		if mes.IsEmpty() != true {
			l.messageChan <- mes.GetMessage()
		}
	}
	close(l.messageChan)
}

func (l *LogReader) generateTailConfig() tail.Config {
	tailConfig := tail.Config{
		Follow: true,
		ReOpen: true,
	}
	if !l.doReadFile {
		tailConfig.Location = &tail.SeekInfo{Offset: 0, Whence: io.SeekEnd}
	}
	return tailConfig
}

func (l *LogReader) Message() chan string {
	return l.messageChan
}

func handleLine(line string) message.Message {
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
