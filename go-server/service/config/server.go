package config

type Server struct {
	Conf *Config
}

func NewServer(config *Config) *Server {
	return &Server{
		Conf: config,
	}
}
