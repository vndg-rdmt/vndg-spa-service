package main

import (
	"silvex/app"
)

func init() {
	app.ConfigureLocalLogger()
}

func main() {
	webserver := Webserver{
		server:      optimizedMultiprocessServer(),
		startupFunc: productionStartup,

		eventCallbackFailedToStart:,
	}
}
