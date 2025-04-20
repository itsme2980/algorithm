package main

import (
	"algorithm/loadbalancer"
	"fmt"
	"time"
)

func main() {
	// Load Balancer
	servers := []string{"http://localhost:8081", "http://localhost:8082"}
	loadBalancer := loadbalancer.NewLoadBalancer(servers)

	// Simulate requests
	for i := 0; i < 10; i++ {
		server := loadBalancer.NextServer()
		fmt.Printf("Request %d is being handled by %s\n", i+1, server)
		time.Sleep(1 * time.Second)
	}
}
