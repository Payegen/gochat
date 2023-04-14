package service

import (
	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"
	"gochat/models"
	"log"
)

func Msghandle(ws *websocket.Conn) {
	for {
		//接受消息
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		msg := models.Message{}
		json.Unmarshal(message, &msg)
		log.Printf("recv: %s", message)
		switch msg.Type {
		case "connect":
			log.Println("connect:" + msg.Type)
			ClientMap[msg.FormId] = ws
			err = ws.WriteMessage(mt, []byte("connect succful"))
			if err != nil {
				log.Println("write:", err)
				break
			}
		case "p2p":
			log.Println("p2p:" + msg.Type)
			Chatp2p(msg.TargetId, []byte(msg.Content), mt)
		}
		//写消息
		err = ws.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func Chatp2p(toname string, msg []byte, mt int) {
	ws, ok := ClientMap[toname]
	if ok {
		ws.WriteMessage(mt, msg)
	}

}
