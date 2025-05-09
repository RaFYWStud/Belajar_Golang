package app

import (
	"golang-tutorial/config"
	"golang-tutorial/internal/server"
	"golang-tutorial/pkg/token"
)

func Start() {
	config.Load()
	token.Load()
	server.Run()
}
