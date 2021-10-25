package rt

import (
	"fmt"
	"net/http"
	"server/route"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     checkOrigin,
}

func checkOrigin(r *http.Request) bool {
	return true
}
func ServeHTTP(c *gin.Context) {
	//判断请求是否为websocket升级请求。
	if websocket.IsWebSocketUpgrade(c.Request) {
		conn, _ := upgrader.Upgrade(c.Writer, c.Request, nil)
		conn.WriteMessage(websocket.TextMessage, []byte("123456798"))
		defer conn.Close()
		for {
			fmt.Println("132456")
			t, c, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("websocket receive fail:", err)
				break
			}
			fmt.Println("789456")
			err = conn.WriteMessage(t, c)
			if err != nil {
				fmt.Println("websocket send fail:", err)
				break
			}
		}
	}
}
func Cors() gin.HandlerFunc {
	return func(res *gin.Context) {
		res.Header("Access-Control-Allow-Origin", "*")
		res.Header("Access-Control-Allow-Headers", "Authorization,Content-Type,Referer,User-Agent")
		res.Header("Access-Control-Allow-Credentials", "true")
		res.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		res.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE")
		res.Header("content-type", "application/json;charset-UTF-8")
		if res.Request.Method == "OPTIONS" {
			res.AbortWithStatus(http.StatusNoContent)
		}
		res.Next()
	}
}
func Initserve() {
	// http.HandleFunc("/ws", ServeHTTP)
	engin := gin.Default()
	engin.Use(Cors())
	engin.POST("/api/register", route.Register)
	engin.POST("/api/login", route.Login)
	engin.POST("/api/sendsms", route.Sendsms)
	engin.POST("/api/test", route.Test)
	engin.GET("/ws", ServeHTTP)
	engin.Run(":5500")
}
