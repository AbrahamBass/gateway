package models

type Route struct {
	Name    string `mapstructure:"name"`
	Context string `mapstructure:"context"`
	Target  string `mapstructure:"target"`
}

type GatewayConfig struct {
	ListenAddr string  `mapstructure:"listenAddr"`
	Routes     []Route `mapstructure:"routes"`
}