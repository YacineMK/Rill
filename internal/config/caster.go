package config

func GetConfig() *Config {
	cfg := LoadConfig("config.local.yaml")
	return cfg
}