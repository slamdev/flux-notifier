package notifier

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

var upgrader = websocket.Upgrader{}

func handleWebsocketRequest(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.WithError(err).Error("failed to upgrade connection")
		return
	}
	defer func() {
		if err := conn.Close(); err != nil {
			logrus.WithError(err).Error("failed to close connection")
		}
	}()

	logrus.Info("websocket connection is open")

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			if strings.Contains(err.Error(), "close 1000") {
				logrus.Info("websocket connection is closed")
			} else {
				logrus.WithError(err).Error("failed to read message")
			}
			break
		}

		logrus.WithField("ws-message", string(message)).Info("received ws message")

		if err := conn.WriteMessage(mt, message); err != nil {
			logrus.WithError(err).Error("failed to write message")
			break
		}
	}
}
