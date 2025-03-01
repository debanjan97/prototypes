package strats

import (
	"math/rand"
)

type Random struct {
	servers []string
}

func NewRandomStrategy(servers []string) *Random {
	return &Random{servers: servers}
}

func (r *Random) NextServer() string {
	return r.servers[rand.Intn(len(r.servers))]
}
