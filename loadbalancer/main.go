package main

import (
	"flag"
	"io"
	"log"
	"net"
	"sync"

	"loadbalancer/strats"
)

var lb *LoadBalancer

var availableServers = []string{
	"localhost:5001",
	"localhost:5002",
	"localhost:5003",
	"localhost:5004",
}

type Stategy interface {
	NextServer() string // returns the endpoint of the next server, depending on the strategy
}

type LoadBalancer struct {
	strategy Stategy
	servers  []string
}

func NewLoadBalancer(strategy Stategy) *LoadBalancer {
	return &LoadBalancer{
		strategy: strategy,
		servers:  availableServers,
	}
}

func init() {
	lb = NewLoadBalancer(strats.NewRandomStrategy(availableServers))
}

func proxy(conn net.Conn) {
	server := lb.strategy.NextServer()
	backendConn, err := net.Dial("tcp", server)
	if err != nil {
		log.Printf("Error connecting to server: %v", err)
		return
	}
	defer conn.Close()
	defer backendConn.Close()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if _, err := io.Copy(conn, backendConn); err != nil {
			log.Printf("Error copying response: %v", err)
		}
	}()

	go func() {
		defer wg.Done()
		if _, err := io.Copy(backendConn, conn); err != nil {
			log.Printf("Error forwarding request: %v", err)
		}
	}()

	wg.Wait()
}

func choseStrategy(strategy string) {
	switch strategy {
	case "round-robin":
		lb = NewLoadBalancer(strats.NewRoundRobinStrategy(availableServers))
	case "random":
		lb = NewLoadBalancer(strats.NewRandomStrategy(availableServers))
	case "first-server":
		lb = NewLoadBalancer(strats.NewFirstServerStrategy(availableServers))
	default:
		log.Fatalf("Invalid strategy: %s", strategy)
	}
}

func main() {
	strategyFlag := flag.String("strategy", "round-robin", "Load balancing strategy (round-robin or random)")
	flag.Parse()

	choseStrategy(*strategyFlag)

	listner, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatal(err)
	}
	defer listner.Close()

	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go proxy(conn)
	}
}
