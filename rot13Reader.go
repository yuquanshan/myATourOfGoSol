package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader)Read(b []byte)(int, error){
	n,err := rot.r.Read(b)
	//var mid byte
	for i:=0; i<n; i++{
		if b[i] >= 65 && b[i] <= 90{
			b[i] = 65+(b[i]-65+13)%26
		}else if b[i] >= 97 && b[i] <= 122{
			b[i] = 97+(b[i]-97+13)%26
		}
	}
	return n,err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	//s := strings.NewReader("bbxlk,dbbbs")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
