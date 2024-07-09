package api

import (
	"log"

	"github.com/ruziba3vich/armiya-gateway/config"
)

type API struct {
	logger *log.Logger
	cfg    *config.Config
}

func New(cfg *config.Config, logger *log.Logger) *API {
	return &API{
		logger: logger,
		cfg:    cfg,
	}
}

func (a *API) RUN() error {
	return nil
}
