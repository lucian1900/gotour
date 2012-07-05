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
    a := "a"[0]
    z := "z"[0]

    n, err := r.r.Read(buf)

    for i, v := range(buf) {
        p := v + 13
        if p > z {
            p = a + (p - z)
        }
        buf[i] = p
    }

    return n, err
}

func main() {
    s := strings.NewReader(
        "Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}