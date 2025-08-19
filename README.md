# Rill

A lightweight real-time streaming platform built with Go that seamlessly handles live videos  
by converting RTMP input streams from encoders like OBS Studio to adaptive HLS output streams  
for smooth web playback across all modern browsers and mobile devices.

![Rill Screenshot](doc/Screenshot%20from%202025-08-19%2015-11-47.png)


## Features

- RTMP to HLS conversion
- JWT authentication 
- Web interface for streaming
- RESTful API
- Cross-platform browser support

## Tech Stack

- **Go 1.21+** - Backend server
- **RTMP** - Stream ingestion
- **HLS** - Web delivery
- **JWT** - Authentication
- **Chi Router** - HTTP routing

## Quick Start

1. **Clone and setup**
```bash
git clone https://github.com/YacineMK/Rill.git
cd Rill
go mod tidy
```

2. **Configure**
```bash
cp config.example.yaml config.local.yaml
# Edit config.local.yaml with your ports and JWT secret
```

3. **Run**
```bash
go run cmd/server/main.go
```

4. **Use**
- Web UI: `http://localhost:8080/static`
- RTMP: `rtmp://localhost:1935`

## How It Works

```
OBS Studio → RTMP → Rill Server → HLS → Web Browser
```

1. Generate stream key in web interface
2. Configure OBS with RTMP URL and token
3. Start streaming
4. Watch in browser using stream ID

