package main

import (
	"context"
	"playground/cpp-bootcamp/api"
	"playground/cpp-bootcamp/api/handler"
	"playground/cpp-bootcamp/config"
	"playground/cpp-bootcamp/pkg/logger"
	"playground/cpp-bootcamp/storage/db"
	"playground/cpp-bootcamp/storage/redis"
)

func main() {

	cfg := config.Load()
	log := logger.NewLogger("mini-project", logger.LevelInfo)
	strg, err := db.NewStorage(context.Background(), cfg)
	redisStrg, err := redis.NewCache(context.Background(), cfg)
	if err != nil {
		return
	}

	h := handler.NewHandler(strg, log, redisStrg)

	r := api.NewServer(h)
	r.Run()
}
