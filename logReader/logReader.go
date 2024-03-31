package logReader

import (
	"github.com/Hazegard/mcsologs/config"
	"github.com/Hazegard/mcsologs/regexParser"
	"github.com/Hazegard/mcsologs/structs/message"
	"github.com/nxadm/tail"
	"io"
)

// LogReader holds the log file parsing implementation
type LogReader struct {
	messageChan chan message.Message
	doReadFile  bool
	tailFile    tail.Tail
	logFile     string
}

// NewLogReader returns a LogReader populated using the configuration
func NewLogReader(c *config.Config) *LogReader {
	return &LogReader{
		messageChan: make(chan message.Message),
		doReadFile:  c.Debug,
		logFile:     c.LogFile,
	}
}

// TailFile reads continuously the log file and wait for a new line to be appended
// When a newline is written, this line is send to the regex parsers,
// If a regex parser matches the line, a struct holding the correspondig message
// is send to the channel
func (l *LogReader) TailFile() {
	t, err := tail.TailFile(l.logFile, l.generateTailConfig())
	if err != nil {
		panic(err)
	}
	for line := range t.Lines {
		mes := handleLine(line.Text)
		if mes.IsEmpty() != true {
			l.messageChan <- mes
		}
	}
	close(l.messageChan)
}

// generateTailConfig return the configuration used by the tail.tailFile method
// If the debug mode is enabled, the log file is read from the beginning of the file
// otherwise, the configuration makes the TailFile waits for new lines
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

// Message returns the channel used to consume the parsed log line
func (l *LogReader) Message() <-chan message.Message {
	return l.messageChan
}

// handleLine parse the line and returns the corresponding Message struct
func handleLine(line string) message.Message {
	var mes message.Message
	mes = regexParser.ParseLogin(line)
	if !mes.IsEmpty() {
		return mes
	}

	mes = regexParser.ParseLogout(line)
	if !mes.IsEmpty() {
		return mes
	}

	mes = regexParser.ParseDeath(line)
	if mes != nil {
		return mes
	}

	mes = regexParser.ParseAdvancement(line)
	if !mes.IsEmpty() {
		return mes
	}

	return mes
}
