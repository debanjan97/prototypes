# TCP Load Balancer Project

A prototype Layer 4 (TCP) load balancer implementation in Go with configurable distribution strategies, along with a test suite to evaluate its performance.

## Project Structure

```
.

├── main.go          # Load balancer implementation
├── README.md        # Load balancer specific documentation
├── servers/
│   ├── app.py          # Flask server implementation
│   ├── Dockerfile      # Server container configuration
│   └── requirements.txt # Server dependencies
├── loadbalancer_test/
│   ├── test_loadbalancer.py  # Load testing script
│   ├── requirements.txt      # Test suite dependencies
│   └── README.md            # Testing suite documentation
└── docker-compose.yml       # Backend servers orchestration
```

## Quick Start

1. Start the backend servers:
```bash
docker compose up --build
```

2. Run the load balancer (in a new terminal):
```bash
cd loadbalancer
go run main.go --strategy round-robin  # or --strategy random
```

3. Run load tests (in a new terminal):
```bash
cd loadbalancer_test
python -m venv venv
source venv/bin/activate  # On Windows: venv\Scripts\activate
pip install -r requirements.txt
python test_loadbalancer.py --requests 1000
```

## Load Balancer Usage

```bash
go run loadbalancer/main.go [flags]

Flags:
  --strategy string   Load balancing strategy (round-robin or random) (default "round-robin")
  --port int         Port to listen on (default 5050)
```

## Testing Suite Usage

```bash
python loadbalancer_test/test_loadbalancer.py [flags]

Flags:
  --requests INT     Number of concurrent requests (default: 1000)
  --port INT        Load balancer port to test (default: 5050)
```

## Backend Servers

The backend servers are simple Flask applications that:
- Listen on ports 5001-5004
- Respond to /health with "I am server X" (X is the server ID)
