// methods/23
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	switch {
		case b >= 'A' && b <= 'Z':
		return ('A' + (b + 13 - 'A') % 26)
		case b >= 'a' && b <= 'z':
		return ('a' + (b + 13 - 'a') % 26)
	}
	return b
}

func (reader rot13Reader) Read(bytes []byte) (int, error) {
	n, err := reader.r.Read(bytes)
	for i := range bytes {
		bytes[i] = rot13(bytes[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
