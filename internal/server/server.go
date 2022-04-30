package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sfninety/iris"
)

var (
	service iris.Service
)

func Init() {
	router := iris.Router{}

	fmt.Println(router.Describe())

	service = router.Serve()

}

func Run() {
	srv, err := iris.Listen(service, ":8000")
	if err != nil {
		panic(err)
	}
	fmt.Printf("listening on %v", srv.GetListener().Addr())

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done

	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	srv.Stop(c)
}
