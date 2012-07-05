package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (r *rot13Reader) Read(buf []byte) (int, error) {
    n, err := r.r.Read(buf)

    for i, v := range(buf) {
        buf[i] = rot13(v)
    }

    return n, err
}

func rot13(l byte) byte {
    a := "a"[0]
    z := "z"[0]

    p := l + 13

    if p > z {
        return a + p - z
    }

    return p
}

func main() {
    s := strings.NewReader(
        "Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}