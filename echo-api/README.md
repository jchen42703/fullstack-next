# Echo API

## Getting Started

Running locally:

```bash
# Linux Only
export $(cat .env | xargs) && go run .
```

Running with Docker:

```bash
# 1. Build
docker build --no-cache -t jchen42703/echo-api:latest .

# 2. If already built, run
docker run -d -p 3000:3000 -e SERVER_URL=0.0.0.0:3000 jchen42703/echo-api
```
