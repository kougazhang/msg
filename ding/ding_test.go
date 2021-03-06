package ding

import (
    "fmt"
    "testing"
)

func TestRequest_Send(t *testing.T) {
    req := Request{
        Url:     "url",
        Keyword: "msg",
    }

    resp, err := req.Send("hello world")
    if err != nil {
        panic(err)
    }

    fmt.Println(resp)
}
