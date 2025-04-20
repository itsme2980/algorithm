package test

import (
	"algorithm/loadbalancer"
	"testing"
	"time"
)

// Test that round-robin cycles correctly through the healthy servers
func TestNextServerRoundRobin(t *testing.T) {
	servers := []string{
		"http://server1",
		"http://server2",
		"http://server3",
	}

	lb := loadbalancer.NewLoadBalancer(servers)

	expected := []string{
		"http://server1",
		"http://server2",
		"http://server3",
		"http://server1",
		"http://server2",
	}

	for i, want := range expected {
		got := lb.NextServer()
		if got != want {
			t.Errorf("RoundRobin %d: expected %s, got %s", i+1, want, got)
		}
	}
}

// Test that NewLoadBalancer initializes with correct healthy servers
func TestLoadBalancerInitialization(t *testing.T) {
	servers := []string{
		"http://s1", "http://s2",
	}
	lb := loadbalancer.NewLoadBalancer(servers)

	if len(lb.NextServer()) == 0 {
		t.Error("Expected at least one healthy server")
	}
}

// Test NextServer returns fallback message if all servers are removed
func TestNextServerEmpty(t *testing.T) {
	lb := loadbalancer.NewLoadBalancer([]string{})

	got := lb.NextServer()
	want := "No healthy servers available"
	if got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
}

// Simulate servers becoming unhealthy after startup (mocking)
func TestNextServerWithHealthCheckDisabled(t *testing.T) {
	// Simulate behavior with manual change (without relying on healthcheck.go)
	servers := []string{
		"http://s1",
		"http://s2",
		"http://s3",
	}
	lb := loadbalancer.NewLoadBalancer(servers)

	// Simulate only s2 is healthy now
	lb.HealthyServers = []string{"http://s2"}

	got := lb.NextServer()
	if got != "http://s2" {
		t.Errorf("Expected s2 as only healthy server, got %s", got)
	}
}

// Test that NextServer returns consistent value with only one healthy server
func TestSingleHealthyServer(t *testing.T) {
	lb := loadbalancer.NewLoadBalancer([]string{"http://onlyone"})
	for i := 0; i < 5; i++ {
		got := lb.NextServer()
		if got != "http://onlyone" {
			t.Errorf("Expected http://onlyone, got %s", got)
		}
	}
}

// Optional: Give time for background health check goroutine to run (if implemented)
func TestBackgroundHealthCheckTick(t *testing.T) {
	lb := loadbalancer.NewLoadBalancer([]string{"http://localhost:9999"})
	time.Sleep(11 * time.Second) // Let health check run once if active

	got := lb.NextServer()
	if got == "" {
		t.Error("Expected fallback message or string, got empty string")
	}
}
