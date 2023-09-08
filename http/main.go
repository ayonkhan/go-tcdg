package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
)

type logWeiter struct{}

func main() {
    resp, err := http.Get("http://google.com")

    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    //bs := make([]byte, 99999)
    //resp.Body.Read(bs)
    //
    //fmt.Println(string(bs))

    lw := logWeiter{}

    io.Copy(lw, resp.Body)
}

func (logWeiter) Write(bs []byte) (int, error) {
    fmt.Println(string(bs))
    fmt.Println("Just wrote this many bytes:", len(bs))

    return len(bs), nil
}
