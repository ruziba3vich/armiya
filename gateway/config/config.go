package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		SoldiersHost   string
		AmmosHost      string
		EquipmentsHost string
		FuelsHost      string
		GatewayHost    string
		AiService      string
	}
)

func (c *Config) Load() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	c.SoldiersHost = os.Getenv("SOLDIERS_HOST")
	c.AmmosHost = os.Getenv("AMMOS_HOST")
	c.EquipmentsHost = os.Getenv(("EQUIPMENTS_HOST"))
	c.FuelsHost = os.Getenv("FUELS_HOST")
	c.GatewayHost = os.Getenv("GATEWAY_HOST")
	c.AiService = os.Getenv("AI_SERVICE")
	return nil
}

func New() (*Config, error) {
	var cnfg Config
	if err := cnfg.Load(); err != nil {
		return nil, err
	}
	return &cnfg, nil
}
