package ding

import (
    "encoding/json"
    "fmt"
    "github.com/kougazhang/msg/lib"
    "github.com/kougazhang/requests"
    "time"
)

type Resp struct {
    Errcode int    `json:"errcode"`
    Errmsg  string `json:"errmsg"`
}

type Request struct {
    Url     string `json:"url"`
    Keyword string `json:"keyword"`
}

type RequestBody struct {
    Msgtype string `json:"msgtype"`
    Text    Text   `json:"text"`
}

type Text struct {
    Content string `json:"content"`
}

func (r Request) Send(msg string) (*Resp, error) {
    if len(r.Keyword) > 0 {
        return r.sendByKeyword(msg)
    }

    return nil, lib.ErrNotImplement
}

func (r Request) sendByKeyword(msg string) (*Resp, error) {
    req := requests.Request{
        URL: r.Url,
        Retry: &requests.Retry{
            Times:    3,
            Interval: time.Second * 1,
        },
    }

    data, err := req.PostJson(RequestBody{
        Msgtype: "text",
        Text:    Text{Content: fmt.Sprintf("[%s]%s", r.Keyword, msg)},
    })
    if err != nil {
        return nil, err
    }

    var resp Resp
    err = json.Unmarshal(data, &resp)
    if err != nil {
        return nil, err
    }
    if resp.Errcode == 0 {
        return nil, nil
    }
    return &resp, err
}
