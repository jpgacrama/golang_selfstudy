package tapewriter

import "io"

type tape struct {
	file io.ReadWriteSeeker
}

func (t *tape) Write(p []byte) (n int, err error) {
	t.file.Seek(0, 0)
	return t.file.Write(p)
}

func (t *tape) SetFile(file io.ReadWriteSeeker) {
	t.file = file
}
