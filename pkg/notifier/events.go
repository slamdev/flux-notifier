package notifier

import (
	"encoding/json"
	"github.com/fluxcd/flux/pkg/event"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func handleEventsRequest(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.WithError(err).Info("failed to read request body")
	}
	var e event.Event
	if err := json.Unmarshal(data, &e); err != nil {
		logrus.WithError(err).WithField("data", string(data)).Error("failed to unmarshal json")
	}
	notify(e)
}
