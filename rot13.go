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
    var a, z byte

    if 'a' <= l && l <= 'z' {
        a = 'a'
        z = 'z'
    } else if 'A' <= l && l <= 'Z' {
        a = 'A'
        z = 'Z'
    } else {
        return l
    }

    p := l + 13

    if p > z {
        return a + p - z - 1
    }

    return p
}

func main() {
    s := strings.NewReader(
        "Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}