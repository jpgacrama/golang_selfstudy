package contextaware

import (
	"io"
)

func NewCancellableReader(rdr io.Reader) io.Reader {
	return rdr
}
