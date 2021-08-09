package limitedReader

import "io"

type LimitedReader struct {
	reader      io.Reader
	remainBytes int64
}

func (lr *LimitedReader) Read(p []byte) (n int, err error) {
	if lr.remainBytes <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > lr.remainBytes {
		p = p[:lr.remainBytes]
	}
	n, err = lr.reader.Read(p)
	lr.remainBytes -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}
