# Fullstack Examples with Ory Authentication

# Architecture

1. Any requests that have to do with auth will be sent to the Ory Kratos server.
2. Any protected backend ednpoints will be checked by Express/Go API using the Kratos Go SDK before allowing access.
   1. All requests to the backend `localhost:3000` will be proxied as `localhost:4000`
3. Any protected frontend endpoints will check the validity of the session cookie with the Kratos JS SDK.
   1. All requests to the frontend `0.0.0.0:4455` will not be proxied. However, you should update:
      1. `UI_BASE_URL` to match the frontend URL in the API `.env`
      2. The urls in `Ory Console \ Custom UI` to match the `UI_BASE_URL`

All of the SDKs interact with the Kratos server (whether it be the self-hosted Docker image or through the hosted Ory network).

## Setup

UI:

```bash
cd ui
yarn dev -- -p 4455
```

Kratos:

```bash
cd deployment
# From Blog Post
# https://www.ory.sh/nextjs-authentication-spa-custom-flows-open-source/
# Just opens kratos @ 4433 and mail server @ 4436
# Current setup has the wrong kratos.yml
docker-compose -f quickstart.yml -f contrib/quickstart/kratos/cloud/quickstart.yml up --build --force-recreate -d

# Use:
docker-compose -f quickstart.yml -f contrib/quickstart/kratos/custom-ui/quickstart.yml up --build --force-recreate -d

# With Postgres
# Make sure that postgres is not already running
# Linux:
# sudo systemctl stop postgresql
# Make sure to URL encode your password:
# https://github.com/ory/kratos/issues/2294
docker-compose -f quickstart.yml -f contrib/quickstart/kratos/with-postgres/quickstart.yml -f quickstart-postgres.yml up --build --force-recreate -d

# From Quick Start
# https://www.ory.sh/docs/kratos/quickstart#clone-ory-kratos-and-run-it-in-docker
# Opens Node.js UI on port 4455
# docker-compose -f quickstart.yml -f quickstart-standalone.yml up --build --force-recreate

# With Postgres
# https://community.ory.sh/t/running-kratos-with-postgres/1908/2
docker-compose -f quickstart.yml -f quickstart-standalone.yml -f quickstart-postgres.yml up --build --force-recreate

# Cleanup
docker-compose -f quickstart.yml -f contrib/quickstart/kratos/cloud/quickstart.yml down
docker-compose -f quickstart.yml rm -fsv
```

## Potential To-Dos

Need to update `contrib/quickstart/kratos/email-password/kratos.yml`?

## To-Dos

- [ ] Practical Go Endpoint setup with Postgres (Sequelize)
  - Don't redirect, just return 401
  - Echo, JSON Middleware
- [ ] Move next-js auth api handlers to Go
  - https://github.com/atreya2011/go-kratos-test
  - https://github.com/ory/hydra/discussions/2873
- [ ] Email SMTP Server Setup
- [ ] Practical UI Endpoint protection with protected dashboard.
