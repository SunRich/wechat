package main

import (
	"fmt"
	"github.com/silenceper/wechat/cache"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/material"
	"net/http"
	"github.com/silenceper/wechat/message"
)

func hello(rw http.ResponseWriter, req *http.Request) {

	//配置微信参数
	config := &wechat.Config{
		AppID:          "your app id",
		AppSecret:      "your app secret",
		Token:          "your token",
		EncodingAESKey: "your encoding aes key",
	}
	wc := wechat.NewWechat(config)

	// 传入request和responseWriter
	server := wc.GetServer(req, rw)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {

		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		return &message.Reply{message.MsgTypeText, text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
}

func main() {
	//配置微信参数
	redis := cache.NewRedis("192.168.118.174:6379","")
	config := &wechat.Config{
		AppID:          "wx235a8df325c1f364",
		AppSecret:      "your app secret",
		Token:          "your token",
		EncodingAESKey: "your encoding aes key",
		Cache:          redis,


	}
	wc := wechat.NewWechat(config)
	m := material.NewMaterial(wc.Context)
	url,err:=m.GetMediaURL("UOPeYV3zne3VdhFO1hG8l8-FBAxyDPtCM7iHNBN0_KqJrfcgMhvruxsu8ybNCGnN")
	fmt.Println(url)
	fmt.Println(err)

	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Printf("start server error , err=%v", err)
	}
}
