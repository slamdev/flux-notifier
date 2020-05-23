package notifier

import (
	"fmt"
	"github.com/fluxcd/flux/pkg/event"
	"github.com/sirupsen/logrus"
	"strings"
)

func notify(e event.Event) {
	switch e.Type {
	case event.EventSync:
		processSyncEvent(e)
		break
	case event.EventCommit:
		processCommitEvent(e)
	default:
		logrus.WithField("event", e).Error("unknown event type")
	}
	logrus.WithField("event-obj", e).
		WithField("event-str", e.String()).
		Debug("received event")
}

func processCommitEvent(e event.Event) {
	metadata := e.Metadata.(*event.CommitEventMetadata)
	logrus.WithField(e.Type, e.ServiceIDs).
		WithField("spec", metadata.Spec.Spec).
		Info("received event")
}

func processSyncEvent(e event.Event) {
	metadata := e.Metadata.(*event.SyncEventMetadata)
	logrus.WithField(e.Type, e.ServiceIDs).
		WithField("errors", formatErrors(metadata.Errors)).
		Info("received event")
}

func formatErrors(errors []event.ResourceError) string {
	list := make([]string, len(errors))
	for i := range errors {
		list[i] = fmt.Sprintf("%s: %s", errors[i].ID, errors[i].Error)
	}
	return strings.Join(list, ",")
}
