package notifier

import (
	"github.com/fluxcd/flux/pkg/event"
	"github.com/sirupsen/logrus"
)

func notify(e event.Event) {
	logrus.WithField("event", e.String()).Info("received event")
	logrus.WithField("event", e).Debug("received event")
}
