package files

import (
	"io"
)

type logWriteCloser []io.WriteCloser

func (l logWriteCloser) Close() (err error) {

	for _, file := range l {
		if err = file.Close(); err != nil {
			return
		}
	}

	return

}

func (l logWriteCloser) Write(p []byte) (n int, err error) {

	for _, v := range l {
		if n, err = v.Write(p); err != nil {
			return
		}
	}

	return

}
