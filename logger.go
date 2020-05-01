package logger

import (
	"fmt"
	"io"
	"os"
	"time"
)

type logger struct {
	outputFile io.Writer
	output     bool
}

func New(filepath string, output bool) *logger {
	f, err := os.OpenFile(filepath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil
	}
	return &logger{
		outputFile: f,
		output:     output,
	}
}

func (l *logger) Println(v interface{}) {
	if l.output {
		fmt.Fprintf(l.outputFile, "%s	%v\n", time.Now().Format(time.RFC1123), v)
	}
}



func (l *logger) Fatal(v interface{}) {
	l.Println(v)
	os.Exit(1)
}
