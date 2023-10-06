package handler

import (
	"playground/cpp-bootcamp/pkg/logger"
	"playground/cpp-bootcamp/storage"
)

type Handler struct {
	storage storage.StorageI
	log     logger.LoggerI
	redis  storage.CacheI
}

func NewHandler(strg storage.StorageI, loger logger.LoggerI,redis  storage.CacheI) *Handler {
	return &Handler{storage: strg, log: loger, redis: redis}
}
