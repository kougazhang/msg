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

type Options struct {
    At
}

type RequestBody struct {
    Msgtype string `json:"msgtype"`
    Text    Text   `json:"text"`
    At      At     `json:"at"`
}

type Text struct {
    Content string `json:"content"`
}

type At struct {
    AtMobiles []string `json:"atMobiles"`
    IsAtAll   bool     `json:"isAtAll"`
}

func (r Request) Send(msg string, options ...func(*Options)) (*Resp, error) {
    optIns := &Options{}
    for _, opt := range options {
        opt(optIns)
    }

    if len(r.Keyword) > 0 {
        return r.sendByKeyword(msg, optIns)
    }

    return nil, lib.ErrNotImplement
}

func (r Request) sendByKeyword(msg string, optIns *Options) (*Resp, error) {
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
        At:      optIns.At,
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
