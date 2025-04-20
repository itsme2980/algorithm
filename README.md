# ⚖️ Round Robin Load Balancer

A concurrent, thread-safe load balancer implemented in Go that distributes incoming requests across multiple servers using the **Round Robin** algorithm with integrated **health checks** and **retry logic**.

---

## 🚀 Features

- 🔄 Round Robin load distribution  
- ❤️ Periodic health checks with retry/backoff  
- 🧵 Safe concurrent server rotation  
- ✅ Unit tests for core logic and behaviors  

---

## 📁 Project Structure

```
round-robin-loadbalancer/
├── loadbalancer/
│   ├── balancer.go        # Core load balancer logic
│   ├── healthcheck.go     # Periodic health checks with retry
│   ├── logger.go          # Logging utilities
│   └── retry.go           # Retry logic
├── test/
│   └── balancer_test.go   # Unit tests
├── main.go                # Sample usage
└── go.mod                 # Module definition
```

---

## ⚙️ Getting Started

### ✅ Prerequisites

- Go **1.20** or higher installed

### 📥 Installation

```bash
git clone https://github.com/itsme2980/round-robin-loadbalancer.git
cd round-robin-loadbalancer
```

### ▶️ Run the Load Balancer

```bash
go run main.go
```

### 🧪 Run Unit Tests

```bash
go test ./test -v
```

---

## 📌 Improvement Opportunities

### 🧪 1. Code Quality & Testing

- Add integration tests using real HTTP servers  
- Add benchmark tests for performance profiling  
- Generate test coverage reports in CI  
- Expand edge case test coverage:
  - Duplicate server handling
  - Rapid health state transitions

### 🧱 2. Architecture Enhancements

- Implement circuit breaker pattern  
- Introduce weighted round-robin logic  

### 🚀 3. Performance Optimizations

- Enable connection pooling  
- Add request queuing logic  
- Introduce caching for frequent requests  

### 🌐 4. API & Management

- REST API endpoints for:
  - Add/remove servers
  - Check health/status
  - Apply runtime configuration  
- Secure API with basic auth/token  

### 📚 5. Documentation

- Add `godoc` comments for all exported methods  
- Generate API documentation  
- Include architecture diagrams  
- Document config options and usage flags  

### 📦 6. Infrastructure & Deployment

- Add Dockerfile for containerization  
- Provide Kubernetes manifests  
- Implement graceful shutdown hooks  
- Add `/healthz` endpoint for the LB itself  
- Export Prometheus metrics or logs  

### 🔐 7. Security Enhancements

- Add TLS/HTTPS support  
- Rate limiting per client/IP  
- ACL-based server filtering  
- Input/request validation  
- Request sanitization  

---

## 🎯 Suggested Implementation Priorities

1. ✅ Complete test coverage  
2. ⚙️ Config management via file/env  
3. 📊 Monitoring & observability  
4. 📖 Documentation and diagrams  
5. 🚀 Performance tuning  
6. 📦 Infra & container support  
7. 🔐 Security hardening