package models

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

// 消息
type Message struct {
	gorm.Model
	FormId   string //发送者
	TargetId string //接受者
	Type     string //消息类型 群聊 私聊 广播
	Media    int    //消息类型 文字 图片 音频
	Content  string //消息内容
	Pic      string
	Url      string
	Desc     string
	Amount   int //其他数字统计
}

func (table *Message) TableName() string {
	return "message"
}

type Client struct {
	Conn          *websocket.Conn //连接
	Addr          string          //客户端地址
	FirstTime     uint64          //首次连接时间
	HeartbeatTime uint64          //心跳时间
	LoginTime     uint64          //登录时间
	Register      chan string     //消息
	//GroupSets     set.Interface   //好友 / 群
}

var ClientMap map[string]*Client = make(map[string]*Client, 0)

func Ws(c *gin.Context) {
	//协议升级
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer func() {
		ws.Close()
		log.Print("ws close;")
	}()
	name := c.Query('username')
	//获取客户端
	currentTime := uint64(time.Now().Unix())
	client := &Client{
		Conn:          ws,
		Addr:          ws.RemoteAddr().String(), //客户端地址
		HeartbeatTime: currentTime,              //心跳时间
		LoginTime:     currentTime,              //登录时间
		Register:      make(chan string, 50),
	}
	client.Register <- name
	ClientMap[name] = client
	Msghandle(ws)
}

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
