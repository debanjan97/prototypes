#!/usr/bin/env python3
"""
Load Balancer Test Script

This script performs concurrent testing of a load balancer by sending multiple HTTP requests
and analyzing the distribution of requests across backend servers. It generates both
statistical output and a visual representation of the load distribution.

Usage:
    python test_loadbalancer.py [--requests N] [--port PORT]

Options:
    --requests N    Number of concurrent requests to make (default: 1000)
    --port PORT    Port of the load balancer (default: 5050)
"""

import aiohttp
import asyncio
import argparse
import matplotlib.pyplot as plt
from collections import Counter
import time

async def make_request(session, url):
    try:
        async with session.get(url) as response:
            text = await response.text()
            # Extract server number from response "I am server X"
            return text.split()[-1]
    except Exception as e:
        print(f"Request failed: {e}")
        return None

async def run_tests(num_requests, port=5050):
    url = f"http://localhost:{port}/health"
    results = []
    
    print(f"Starting {num_requests} concurrent requests to {url}")
    start_time = time.time()

    # Create a session for connection pooling
    async with aiohttp.ClientSession() as session:
        # Create list of tasks for concurrent execution
        tasks = [make_request(session, url) for _ in range(num_requests)]
        # Execute all requests concurrently
        results = await asyncio.gather(*tasks)
    
    end_time = time.time()
    
    # Filter out None values (failed requests)
    results = [r for r in results if r is not None]
    
    # Count occurrences of each server
    counter = Counter(results)
    
    # Print statistics
    print(f"\nTest completed in {end_time - start_time:.2f} seconds")
    print("\nDistribution of requests:")
    for server_id, count in sorted(counter.items()):
        print(f"Server {server_id}: {count} requests ({count/len(results)*100:.1f}%)")
    
    # Create bar plot
    plt.figure(figsize=(10, 6))
    servers = sorted(counter.keys())
    counts = [counter[server] for server in servers]
    
    # Create bar plot with server IDs as categorical x-axis
    x_pos = range(len(servers))
    plt.bar(x_pos, counts)
    plt.xticks(x_pos, [f'Server {s}' for s in servers])
    
    plt.title(f'Load Distribution Across Servers\n({num_requests} total requests)')
    plt.xlabel('Server')
    plt.ylabel('Number of Requests')
    
    # Add value labels on top of each bar
    for i, count in enumerate(counts):
        plt.text(i, count, str(count), ha='center', va='bottom')
    
    # Save the plot
    plt.savefig('load_distribution.png')
    print("\nPlot saved as 'load_distribution.png'")

def main():
    parser = argparse.ArgumentParser(description='Test load balancer distribution')
    parser.add_argument('--requests', type=int, default=1000,
                      help='Number of concurrent requests to make (default: 1000)')
    parser.add_argument('--port', type=int, default=5050,
                      help='Port of the load balancer (default: 5050)')
    
    args = parser.parse_args()
    
    # Run the async test
    asyncio.run(run_tests(args.requests, args.port))

if __name__ == "__main__":
    main() 