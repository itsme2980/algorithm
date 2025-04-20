package loadbalancer

import (
	"sync"
)

type LoadBalancer struct {
	servers        []string
	HealthyServers []string
	currentIndex   int
	mu             sync.Mutex
}

func NewLoadBalancer(servers []string) *LoadBalancer {
	lb := &LoadBalancer{
		servers:        servers,
		HealthyServers: servers,
	}

	go lb.startHealthCheck()
	return lb
}

func (lb *LoadBalancer) NextServer() string {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	if len(lb.HealthyServers) == 0 {
		return "No healthy servers available"
	}

	server := lb.HealthyServers[lb.currentIndex]
	lb.currentIndex = (lb.currentIndex + 1) % len(lb.HealthyServers)
	return server
}
