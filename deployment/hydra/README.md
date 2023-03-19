# Ory Hydra Deployment

With Docker Compose for simplicity purposes.

The general architecture is:

- UI
- API
- Hydra as proxy between UI and API

## Getting Started

Basic getting started example:

```bash
docker-compose -f quickstart.yml \
-f quickstart-postgres.yml \
up --build
```
