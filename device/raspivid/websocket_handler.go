package raspivid

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"pimonitor/device/servo"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upGrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		}}
)

type webSocketHandler struct {
	opts    Options
	unicast bool
	mutex   sync.Mutex
}

func NewWebSocketHandler(opts Options) *webSocketHandler {
	return &webSocketHandler{
		opts: opts,
	}
}

func (w *webSocketHandler) Handler(c *gin.Context) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	if w.unicast {
		log.Println("The resource is already occupied")
		return
	}

	// Upgrade
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Upgrade websocket request error: %v \n", err)
		return
	}
	go w.streaming(conn)
}

func (w *webSocketHandler) streaming(conn *websocket.Conn) {
	w.unicast = true
	log.Printf("New websocket connection from %s success.\n", conn.RemoteAddr())

	ctx, cancel := context.WithCancel(context.Background())
	getter := make(chan []byte, 10e2)
	go SendVideoStream(ctx, w.opts, getter)

	defer func() {
		log.Printf("Websocket connection from %s closed\n", conn.RemoteAddr())
		cancel()
		conn.Close()
		w.unicast = false
	}()

	wsClosed := make(chan struct{})
	go func() {
		defer close(wsClosed)
		for {
			_, p, err := conn.ReadMessage()
			if err != nil {
				return
			}
			param := &ServoCmd{}
			if err := json.Unmarshal(p, param); err != nil {
				continue
			}
			switch param.Direction {
			case "left":
				servo.X.Reduce(param.Value)
			case "right":
				servo.X.Add(param.Value)
			case "up":
				servo.Y.Reduce(param.Value)
			case "down":
				servo.Y.Add(param.Value)
			}
		}
	}()

	var (
		count int = w.opts.FPS
		timer     = time.NewTicker(time.Second)
	)
Loop:
	for {
		select {
		case data, ok := <-getter:
			if !ok {
				break Loop
			}

			if count <= 0 {
				continue
			}
			count--
			if err := conn.WriteMessage(websocket.BinaryMessage, data); err != nil {
				log.Printf("Write websocket conn %s error: %v\n", conn.RemoteAddr(), err)
				break Loop
			}
		case <-timer.C:
			count = 30
		case <-wsClosed:
			break Loop
		}
	}
}

type ServoCmd struct {
	Direction string `json:"direction" `
	Value     uint8  `json:"value"`
}
