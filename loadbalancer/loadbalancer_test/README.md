# Load Balancer Test Suite

This test suite provides tools to evaluate the performance and distribution characteristics of a load balancer implementation.

## Setup

1. Create and activate a virtual environment (recommended):
```bash
python -m venv venv
source venv/bin/activate  # On Windows: venv\Scripts\activate
```

2. Install dependencies:
```bash
pip install -r requirements.txt
```

## Usage

The test script supports various command-line arguments to customize the test:

```bash
# Run with default settings (1000 requests to port 5050)
python test_loadbalancer.py

# Run with custom number of requests
python test_loadbalancer.py --requests 5000

# Run with custom port
python test_loadbalancer.py --port 5050

# Run with both custom requests and port
python test_loadbalancer.py --requests 5000 --port 5050
```

## Output

The script provides two types of output:

1. Console output with:
   - Total execution time
   - Request distribution statistics
   - Percentage of requests handled by each server

2. Visual output:
   - Generates 'load_distribution.png'
   - Bar chart showing request distribution
   - Labels showing exact count for each server

## Requirements

All dependencies are listed in requirements.txt with exact versions to ensure reproducibility. Main dependencies include:
- aiohttp: For concurrent HTTP requests
- matplotlib: For visualization
- Various supporting packages 