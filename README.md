<p align="center">
  <img src="./assets/Proxytron.jpg" alt="ProxyTron Logo" width="400"/>
</p>

<p align="center">
  A fast, lightweight HTTP proxy server with Redis-based caching, built with Go.
</p>

---

## Features

- HTTP proxy with response caching
- Redis-based cache integration
- Built with Go for high performance and concurrency
- Dockerized for easy local development
- API key-based authentication (planned)
- Logging and observability (planned)
- Web dashboard (planned)

---

## Tech Stack

- **Language**: Go (Golang)
- **Cache**: Redis
- **Containers**: Docker, Docker Compose
- **Dependencies**: `net/http`, `github.com/redis/go-redis/v9`

---

## Getting Started

### Prerequisites

- Go 1.20 or above
- Docker & Docker Compose installed

### Clone the Repository

```bash
git clone https://github.com/<your-username>/proxytron.git
cd proxytron
```

### Run Locally

```bash
go run main.go
```

This runs the proxy server on:

```
http://localhost:8080
```

Make sure Redis is running locally at `localhost:6379` or set the `REDIS_ADDR` environment variable.

### Run with Docker Compose

```bash
docker-compose up --build
```

This setup runs:
- The ProxyTron server on `localhost:8080`
- A Redis container for caching

---

## Example Request

Once the server is running, you can make proxy requests like this:

```bash
curl "http://localhost:8080/?url=https://example.com"
```

- The **first request** will be fetched from the origin and cached.
- **Subsequent requests** will be served from Redis cache.

---

## Roadmap

- [x] Basic HTTP proxy
- [x] LRU caching (in-memory)
- [x] Redis cache integration
- [ ] API key authentication middleware
- [ ] Logging and metrics
- [ ] Web UI for cache stats and config

---

## License

MIT License  
© 2025 [Anurag Malasi](https://github.com/<your-username>)

---

## Contributing

Contributions are welcome. Feel free to fork the repo, raise issues, or submit pull requests.

---

## Contact

For bugs, feature requests, or suggestions — open an issue on GitHub or connect via your GitHub profile.
