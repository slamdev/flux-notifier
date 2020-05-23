package notifier

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func handleEventsRequest(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.WithError(err).Info("failed to read request body")
	}
	event := make(map[string]interface{})
	if err := json.Unmarshal(data, &event); err != nil {
		logrus.WithError(err).WithField("data", string(data)).Error("failed to unmarshal json")
	}
	logrus.WithField("event", event).Info("received event")
}
