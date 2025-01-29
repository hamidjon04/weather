package handler

import (
	"log/slog"
	"weather/service"
)

type Handler struct{
	Service service.Service
	Log *slog.Logger
}

func NewHandler(service service.Service, log *slog.Logger)Handler{
	return Handler{
		Service: service,
		Log: log,
	}
}