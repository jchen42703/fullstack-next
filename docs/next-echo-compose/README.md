# Deploy Next.js UI + Echo Go API + Ory Kratos + Postgres with Docker-Compose and NGINX <!-- omit in toc -->

This document describes the deployment in [`deployment/compose/postgres-nginx`](../../deployment/compose/postgres-nginx/).

## Table Of Contents <!-- omit in toc -->

- [Getting Started](#getting-started)
- [Architecture](#architecture)
- [Deploy to Production](#deploy-to-production)
- [Customization](#customization)

## Getting Started

All of the commands below assume that your current working directory is `deployment/compose`. Adjust your filepaths to the `postgres-nginx/docker-compose.yml` accordingly.

Before starting, make sure that you:

1. Have `docker` and `docker-compose` installed.
2. Have `postgres` installed (TODO: Double check this)
3. Make sure that the following ports are open:
   1. `4000`: For the NGINX entrypoint
   2. `4433`, `4434`, `4436`: For Kratos' Public API, Kratos' Admin API, and Kratos' mail server.
   3. `4445`: For the custom UI
   4. `3000`: For the backend.

The NGINX, UI, and backend ports can be easily customized for your needs. This will be covered in [Customization](#customization).

```bash
# Make sure that postgres is not already running locally.
# Linux:
# sudo systemctl stop postgresql
# Make sure to URL encode your password:
# https://github.com/ory/kratos/issues/2294
docker-compose -f postgres-nginx/docker-compose.yml up --build --force-recreate -d
```

You should see something like:

```bash
Recreating postgres-nginx_kratos-migrate_1 ... done
Recreating postgres-nginx_mailslurper_1    ... done
Recreating postgres-nginx_postgresd_1      ... done
Recreating postgres-nginx_backend_1        ... done
Recreating postgres-nginx_kratos_1         ... done
Recreating postgres-nginx_ui_1             ... done
Recreating postgres-nginx_proxy_1          ... done
```

To shut down the containers, simply do:

```bash
docker-compose -f postgres-nginx/docker-compose.yml down
```

**Optional:** Clean up volumes with:

```bash
docker-compose -f postgres-nginx/docker-compose.yml rm -fsv
```

## Architecture

## Deploy to Production

## Customization
