package api

import (
	"time"

	"github.com/Gimulator/Gimulator/object"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 5 * time.Second

	// Send pings to peer with this period.
	pingPeriod = time.Second * 3
)

type client struct {
	username string
	ch       chan *object.Object
	token    string
	log      *logrus.Entry
}

func NewClient(username string, token string) *client {
	return &client{
		username: username,
		token:    token,
		ch:       make(chan *object.Object),
		log:      logrus.WithField("Entity", "client"),
	}
}

func (c *client) GetChan() chan *object.Object {
	return c.ch
}

func (c *client) GetToken() string {
	return c.token
}

func (c *client) Reconcile(conn *websocket.Conn) {
	c.log.Info("Start to write")
	defer c.log.Debug("End of writing to the connection")

	var err error
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		conn.Close()
	}()
	for {
		select {
		case obj, ok := <-c.ch:
			if !ok {
				conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			conn.SetWriteDeadline(time.Now().Add(writeWait))
			err = conn.WriteJSON(obj)
			if err != nil {
				c.log.WithError(err).Error("Can not write json to connection")
				return
			}
		case <-ticker.C:
			conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
