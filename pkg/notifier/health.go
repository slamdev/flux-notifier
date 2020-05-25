package notifier

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

func handleHealthRequest(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintf(w, "OK"); err != nil {
		logrus.WithError(err).Error("failed to write response")
	}
}
