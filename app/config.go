package app

// A default application config
var Config *ApplicationConfig

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

type ApplicationConfig struct {
	Server  ServerConfig
	Client  ClientConfig
	Logs    LogsConfig
	Limiter LimiterConfig
}

// Application parts configurations
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
