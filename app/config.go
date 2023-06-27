package app

/*
Holds a defailt application config
*/
var Config *ApplicationConfig

type ApplicationConfig struct {
	Server  ServerConfig
	Client  ClientConfig
	Logs    LogsConfig
	Limiter LimiterConfig
}

// ===========================

type ServerConfig struct {
	Addr string
	Port string
}

type ClientConfig struct {
	Directory string
	Document  string
}

type LimiterConfig struct {
	Enable   bool
	Retries  int
	JailTime int
}

type LogsConfig struct {
	Addr string
	Port string
}

func init() {
	Config = &ApplicationConfig{
		Server: ServerConfig{
			Addr: "127.0.0.1",
			Port: "9999",
		},
		Client: ClientConfig{
			Directory: "./client-app",
			Document:  "index.html",
		},
		Limiter: LimiterConfig{
			Retries:  10,
			JailTime: 15,
		},
	}
}
