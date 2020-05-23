package notifier

import (
	"context"
	"github.com/sirupsen/logrus"
)

func Listen(ctx context.Context, port int) error {
	logrus.WithField("port", port).Info("listening for events")
	logrus.Info("shutting down listener")
	return nil
}
