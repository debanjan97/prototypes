package strats

import "sync"

type RoundRobinStrategy struct {
	servers      []string
	currentIndex int
	mu           sync.Mutex
}

func NewRoundRobinStrategy(servers []string) *RoundRobinStrategy {
	return &RoundRobinStrategy{
		servers:      servers,
		currentIndex: 0,
	}
}

func (s *RoundRobinStrategy) NextServer() string {
	s.mu.Lock()
	defer s.mu.Unlock()

	server := s.servers[s.currentIndex]
	s.currentIndex = (s.currentIndex + 1) % len(s.servers)
	return server
}
