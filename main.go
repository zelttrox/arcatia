package main

import (
	"log"
	"main/src/engine"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	var e engine.Engine

	e.Init()
	e.Load()
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	e.Run()
	e.Unload()
	e.Close()
}
