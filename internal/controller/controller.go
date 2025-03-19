package controller

import (
	"github.com/AndreySurzhan/messy-room-api/internal/config"
)

type Service interface {
	GetRoomCleanlinessStatus(image string) (string, error)
}

type Controller struct {
	service Service
	config  *config.Config
}

func New(service Service, cfg *config.Config) *Controller {
	return &Controller{
		service: service,
		config:  cfg,
	}
}
