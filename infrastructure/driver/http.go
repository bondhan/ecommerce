package driver

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunHttpServer(logger *logrus.Logger, port string, r *chi.Mux) {
	server := &http.Server{Addr: ":" + port, Handler: r}
	go func() {
		logger.Info("application started at port:", port)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			logger.Error(err)
			return
		}
	}()

	// Setting up a channel to capture system signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	//wait forever until 1 of signals above are received
	<-stop

	// send warning that we are closing
	logger.Warnf("got signal: %v, closing DB connection gracefully", stop)

	// wait 5 second in background while server is trying to shut down
	ctxKafka, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//try to shut down the server
	logger.Warn("shutting down http server")
	if err := server.Shutdown(ctxKafka); err != nil {
		logger.Error(err)
	}
}
