package api

import (
	"IAGE-test-task/services/task1"
	"IAGE-test-task/services/task2"
	"IAGE-test-task/services/task3"
	"flag"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
)

type App struct {
	listener net.Listener
	server   *http.Server
}

func (a *App) Start() error {
	return a.server.Serve(a.listener)
}

func NewDB() *redis.Client {
	redisHost := flag.String("host","localhost","a string")
	redisPort := flag.String("port", "6060", "string")
	flag.Parse()
	rdb := redis.NewClient(&redis.Options{
		Addr: *redisHost + ":" + *redisPort,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	log.Println(*redisHost, *redisPort)
	log.Println(rdb)
	return rdb
}

func NewListener() (net.Listener, error) {
	return net.Listen("tcp", ":8080")
}

func NewRouter(task1Handlers task1.HandlersHTTP, task2Handlers task2.HandlersHTTP, task3Handlers task3.HandlersHTTP) http.Handler {
	r := mux.NewRouter().StrictSlash(true)
	r.Handle("/task1", task1Handlers.Find).Methods("POST")
	r.Handle("/task2", task2Handlers.Find).Methods("POST")
	r.Handle("/task3", task3Handlers.Find).Methods("POST")
	return r
}

func NewApp(listener net.Listener, handler http.Handler) *App {
	return &App{
		listener: listener,
		server:   &http.Server{Handler: handler},
	}
}