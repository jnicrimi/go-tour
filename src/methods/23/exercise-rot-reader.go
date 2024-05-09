package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(p []byte) (n int, e error) {
	n, e = r.r.Read(p)
	for i := range p {
		p[i] = rot13(p[i])
	}
	return
}

func rot13(n byte) byte {
	switch {
	case ('A' <= n && n <= 'Z'):
		return (n-'A'+13)%26 + 'A'
	case ('a' <= n && n <= 'z'):
		return (n-'a'+13)%26 + 'a'
	default:
		return n
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	_, err := io.Copy(os.Stdout, &r)
	if err != nil {
		panic(err)
	}
}
