package app

import (
	"notify-app/handlers"
)

func initRouters() {
	router.GET("/status", handlers.GetStatus)
}
