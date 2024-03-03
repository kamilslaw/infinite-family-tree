package main

import (
	"context"
	"github.com/kamilslaw/infinite-family-tree/logger"
)

func main() {
	log := logger.InitLogger()
	defer log.Sync()

	log.Info(context.Background(), "Hello world!", logger.Fields{"name": "aaa", "value": 7})
}
