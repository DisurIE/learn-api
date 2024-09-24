package main

import (
	"fmt"
	"learn-api/internal/config"
)

func MustLoad() {

	cfg := config.Load()

	fmt.Println(cfg)

	// TODO: init logger: slog

	// TODO: init storage: sqlite

	// TODO: init router: chi, render

	// TODO: run server
}
