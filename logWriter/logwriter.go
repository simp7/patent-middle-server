package logWriter

import (
	"os"
)

type LogWriter struct {
	files []*os.File
}

func New(logFiles ...*os.File) *LogWriter {

	files := make([]*os.File, 1)
	files[0] = os.Stdout

	files = append(files, logFiles...)

	return &LogWriter{files}

}

func (l *LogWriter) Close() (err error) {

	for _, file := range l.files {
		if err = file.Close(); err != nil {
			return
		}
	}

	return

}

func (l *LogWriter) Write(p []byte) (n int, err error) {

	for _, v := range l.files {
		if n, err = v.Write(p); err != nil {
			return
		}
	}

	return

}
