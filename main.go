package main

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/wings-software/dlite/bijou"
	"github.com/wings-software/dlite/delegate"
	"github.com/wings-software/dlite/poller"
	"github.com/wings-software/dlite/router"
	"github.com/wings-software/dlite/task"
)

type BijouTask struct{}

func (t *BijouTask) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	logrus.Info("reached handler")
}

func main() {
	routes := map[string]task.Handler{
		"SETUP": &BijouTask{},
	}
	router := router.NewRouter(routes)
	// Create a delegate client
	client := delegate.New("https://localhost:9090", "kmpySmUISimoRrJL6NL73w", "2f6b0988b6fb3370073c3d0505baee59", true, "")

	// The poller needs a client that interacts with the task management system and a router to route the tasks
	poller := poller.New("kmpySmUISimoRrJL6NL73w", "2f6b0988b6fb3370073c3d0505baee59", "dlite-xingchi", []string{"k8s-runner"}, client, router)

	// // Register the poller
	ctx := context.Background()
	info, _ := poller.Register(ctx)
	ctx = context.WithValue(ctx, "delegateId", info.ID)
	ctx = context.WithValue(ctx, "delegateName", info.Name)
	logrus.Info(info)


	// Start polling for bijou events
	eventsServer := bijou.New(client)
	eventsServer.PollRunnerEvents(ctx, 3, info.ID, time.Second*10)


    // Just to keep it running
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		logrus.Info(w, "Hello!")
	})

	logrus.Info("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logrus.Fatal(err)
	}

}
