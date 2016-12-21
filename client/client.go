// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	//	"bytes"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	//	"time"

	"IceAndFire/protomsg"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	defer c.Close()

	go func() {
		defer c.Close()
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	cmd := make(chan int32)
	go func() {
		for {
			var input int32
			fmt.Scanf("%d", &input)
			cmd <- input
		}
	}()

	for {
		select {
		case i := <-cmd:
			switch i {
			case 1:
				var msg protomsg.CSMessage
				msg.Head = &protomsg.Head{protomsg.MSG_ID_MSG_ID_TEST}
				msg.Chat = &protomsg.Chat{Send: "Chat message"}
				bmsg, _ := proto.Marshal(&msg)
				err := c.WriteMessage(websocket.BinaryMessage, bmsg)
				if err != nil {
					log.Println("write:", err)
					return
				}
			default:
				log.Println("unknow message")
			}
		case <-interrupt:
			log.Println("interrupt")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			return
		}
	}
}
