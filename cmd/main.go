package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/tomasmik/winter-is-coming/server"

	"github.com/fln/pprotect"
	"github.com/sirupsen/logrus"
	"github.com/vrischmann/envconfig"
)

// this conf might be overkill, but it interacts
// nicely with the env dump made in the Makefile.
var conf struct {
	Port int `envconfig:"default=8081"`
}

func main() {
	if err := envconfig.InitWithOptions(&conf, envconfig.Options{
		Prefix:          "WIC",
		AllowUnexported: true,
	}); err != nil {
		logrus.WithError(err).Fatal("parsing environment variables")
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		logrus.WithError(err).Fatal("failed to start a server")
	}

	server := server.New(l)
	var wg sync.WaitGroup
	wg.Add(1)
	go pprotect.CallLoop(func() {
		defer wg.Done()
		server.Run()
	}, 1*time.Second, panicHandler)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	server.Stop()
	wg.Wait()
}

func panicHandler(val interface{}, stack []byte) {
	logrus.WithFields(logrus.Fields{
		"value": val,
		"stack": string(stack),
	}).Error("thread panic")
}
