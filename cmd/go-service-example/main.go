// Package main provides main service.
package main

import (
	"github.com/powerman/narada-go/narada/bootstrap"

	"net/http"
	"strings"

	"github.com/powerman/narada-go/narada"

	_ "github.com/powerman/go-service-narada-example/api/eventsvc"
	_ "github.com/powerman/go-service-narada-example/api/rpcsvc"
	_ "github.com/powerman/go-service-narada-example/api/websvc"
)

var log = narada.NewLog("")
var listen string

func init() {
	listen = narada.GetConfigLine("listen")
	if strings.Index(listen, ":") == -1 {
		log.Fatal("please setup config/listen")
	}
}

func main() {
	bootstrap.Unlock()
	log.NOTICE("Listening on %s", listen)
	log.Fatal(http.ListenAndServe(listen, nil))
}
