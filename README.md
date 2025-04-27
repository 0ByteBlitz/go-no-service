# Go-No-Service Proxy & Local IP Display

## Overview
This project contains three main components:
1. **Go Proxy Server**: A reverse proxy that forwards requests to a target URL and performs load testing by making concurrent requests through a proxy.
2. **Local Server (Python)**: A Python-based server that serves the local machine’s IP address at `/get-ip`.
3. **HTML Interface**: A simple web page that dynamically fetches and displays the local server's IP.

## Usage

### Proxy Server
To run the Go proxy:
```bash
go run main.go <proxyURL> <targetURL> <numRequests>
