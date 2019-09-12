package websocket

import (
	"battery-anlysis-platform/app/main/service"
	"battery-anlysis-platform/pkg/jd"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func GetTaskList(c *gin.Context) {
	conn, err := upgradeHttpConn(c.Writer, c.Request)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	for {
		data, err := service.GetTaskList()
		code, msg := jd.HandleError(err)
		res := jd.Build(code, msg, data)
		if err := conn.WriteJSON(res); err != nil {
			log.Println(err)
			break
		}
		time.Sleep(time.Second * 3)
	}
}
