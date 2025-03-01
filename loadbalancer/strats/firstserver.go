package strats

type FirstServer struct {
	servers []string
}

func (s *FirstServer) NextServer() string {
	return s.servers[0]
}

func NewFirstServerStrategy(servers []string) *FirstServer {
	return &FirstServer{servers: servers}
}
