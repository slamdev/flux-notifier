package notifier

import (
	"context"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func Listen(ctx context.Context, port int) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		oscall := <-c
		logrus.WithField("signal", oscall).Info("system call")
		cancel()
	}()

	return serve(ctx, port)
}

func serve(ctx context.Context, port int) (err error) {
	mux := http.NewServeMux()
	mux.Handle("/health", http.HandlerFunc(handleHealthRequest))
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/v6/events", http.HandlerFunc(handleEventsRequest))
	mux.Handle("/", http.HandlerFunc(handleWebsocketRequest))

	srv := &http.Server{
		Addr:    ":" + strconv.FormatInt(int64(port), 10),
		Handler: mux,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.WithError(err).Fatal("fail to listen")
		}
	}()

	logrus.WithField("port", port).Info("listening for events")

	<-ctx.Done()

	logrus.Info("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(ctxShutDown); err != nil {
		logrus.WithError(err).Fatal("fail to shutdown")
	}

	logrus.Info("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	logrus.Info("shutting down listener")

	return nil
}
