# Deployment

Deploying the kratos infrastructure.

## Docker Compose

**Quick start with Postgres + Nginx + UI + Backend:**

```bash
# Make sure that postgres is not already running locally.
# Linux:
# sudo systemctl stop postgresql
# Make sure to URL encode your password:
# https://github.com/ory/kratos/issues/2294
docker-compose -f postgres-nginx/docker-compose.yml up --build --force-recreate -d
```

To shut down the containers, simply do:

```bash
docker-compose -f postgres-nginx/docker-compose.yml down
```

**Optional:** Clean up volumes with:

```bash
docker-compose -f postgres-nginx/docker-compose.yml rm -fsv
```

## Older Commands

Starter with postgres:

```bash
# Make sure that postgres is not already running locally.
# Linux:
# sudo systemctl stop postgresql
# Make sure to URL encode your password:
# https://github.com/ory/kratos/issues/2294
# This will create the Kratos API, a Postgres container with a mounted volume, and the mail server.
docker-compose -f custom-postgres.yml up --build --force-recreate -d

# Shutdown with:
docker-compose -f custom-postgres.yml down
```

You can verify your Postgres deployment by:

1. `docker ps` to get the container id of the postgres container.
2. `docker exec -it CONTAINER_ID bash`
3. `psql -U kratos`
4. `SELECT * FROM identities;`
   1. Displays all users. Should show users from previous times you've started the containers.

**Optional:** Clean up volumes with:

```bash
docker-compose -f custom-postgres.yml rm -fsv
```

Misc commands:

```bash
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

## Notes

There are some nuanced aspects of deploying with Docker.

1. Serve servers under `0.0.0.0`.
   1. This is so that everything under the subnet can access the server. This is important for Docker deployments because we will not know what the exact IP will be.
2. Use container names to replace hosts!
   1. For example, in `proxy/nginx.conf`, I used `client` and `backend` to abstract their host IP addresses. This is because in a Docker subnet, the actual IPs of each container will vary and will rarely be `localhost`.
   2. You can also see this when I specify the Kratos Public API Url with `http://kratos:4433`, where `kratos` represents the host IP of the Kratos API container.
3. `docker-compose up` doesn't seem to rebuild with the updated `nginx` config? Not too sure about this one.
