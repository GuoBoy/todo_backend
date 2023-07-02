package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"todo_backend/consts/ws_consts"
	"todo_backend/models"
)

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandle(ctx *gin.Context) {
	conn, err := upgrade.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	for {
		var resp CryptForm
		if err = conn.ReadJSON(&resp); err != nil {
			fmt.Println("解析请求失败: ", err)
			break
		}
		msg, _ := DecryptWsRequest[models.WsMessageModel](resp)
		switch msg.Type {
		case ws_consts.Ping:
			if err = conn.WriteJSON(EncryptToForm(gin.H{"data": "pong"})); err != nil {
				fmt.Println("发送pongerr")
				break
			}
			//case ws_consts.Groups:

			continue
		default:
			continue
		}
	}
}
