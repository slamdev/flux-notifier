package notifier

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func handleEventsRequest(w http.ResponseWriter, r *http.Request) {
	event, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.WithError(err).Info("failed to read request body")
	}
	logrus.WithField("event", string(event)).Info("received event")
}
