package main

import (
	"github.com/shyamz-22/syslog-client/logger"
	"log"
	"log/syslog"
	"net/http"

	"github.com/shyamz-22/syslog-client/api"
)

func main() {

	sysLogger, err := syslog.Dial("tcp", "localhost:2360", syslog.LOG_INFO, "testtag")
	if err != nil {
		log.Fatal("failed to dial syslog")
	}

	l := logger.NewLogger(logger.Debug)
	l.SetSysLogger(sysLogger)
	apis := api.NewServer(l)

	mux := http.NewServeMux()
	mux.HandleFunc("/", apis.Home)
	mux.HandleFunc("/ping", apis.Health)
	mux.HandleFunc("/fatal", apis.Fatal)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
