package config

import "os"

func APISecret() string {
	return os.Getenv("TRADING_DISCORD_API_KEY")
}