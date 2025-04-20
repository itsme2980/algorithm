package loadbalancer

import (
	"net/http"
	"sync"
	"time"
)

var (
	retryCount    = 3
	checkInterval = 10 * time.Second
	debounceMap   = make(map[string]int)
	debounceMutex sync.Mutex
)

func (lb *LoadBalancer) startHealthCheck() {
	ticker := time.NewTicker(checkInterval)
	for range ticker.C {
		lb.checkServers()
	}
}

func (lb *LoadBalancer) checkServers() {
	var healthy []string

	for _, server := range lb.servers {
		if isServerHealthyWithRetry(server) {
			healthy = append(healthy, server)
		}
	}

	lb.mu.Lock()
	lb.HealthyServers = healthy
	lb.currentIndex = 0
	lb.mu.Unlock()

	LogInfo("Health Check Done. Healthy Servers: %v", healthy)
}

func isServerHealthyWithRetry(server string) bool {
	debounceMutex.Lock()
	count := debounceMap[server]
	debounceMutex.Unlock()

	for i := 0; i < retryCount; i++ {
		if isServerHealthy(server) {
			if count > 0 {
				LogInfo("Server %s recovered after %d retries", server, count)
			}
			debounceMutex.Lock()
			debounceMap[server] = 0
			debounceMutex.Unlock()
			return true
		}
		time.Sleep(1 * time.Second)
	}

	debounceMutex.Lock()
	debounceMap[server]++
	debounceMutex.Unlock()
	LogWarn("Server %s marked unhealthy after %d failed attempts", server, retryCount)
	return false
}

func isServerHealthy(url string) bool {
	client := http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get(url + "/health")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}
