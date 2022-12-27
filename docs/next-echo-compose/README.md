# Deploy Next.js UI + Echo Go API + Ory Kratos + Postgres with Docker-Compose and NGINX <!-- omit in toc -->

This document describes the deployment in [`deployment/compose/postgres-nginx`](../../deployment/compose/postgres-nginx/).

## Table Of Contents <!-- omit in toc -->

- [Getting Started](#getting-started)
- [Architecture](#architecture)
- [Deploy to Production](#deploy-to-production)
  - [Cookie Secret](#cookie-secret)
- [Kratos Logging](#kratos-logging)
  - [Postgres Secret](#postgres-secret)
  - [Kratos Redirect URLs](#kratos-redirect-urls)
  - [Email Server Setup](#email-server-setup)
  - [NGINX HTTPS + Domain Deployment](#nginx-https--domain-deployment)
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

This section covers various aspects that you must cover while deploying this example to production. **Make sure to always do these changes prior to `docker-compose` OR `docker-compose down && docker-compose up` again to ensure that the new configurations take place!**

### Cookie Secret

In `postgres-nginx/kratos/kratos.yml`, look for:

```yaml
secrets:
  cookie:
    - PLEASE-CHANGE-ME-I-AM-VERY-INSECURE
```

Replace the string with the result of:

```bash
openssl rand -base64 32
```

## Kratos Logging

In `postgres-nginx/kratos/kratos.yml`, look for:

```yaml
log:
  level: info
  format: text
```

Change to:

```yaml
log:
  level: info
  format: text
  leak_sensitive_values: false
```

### Postgres Secret

You should change the `secret` in the `docker-compose.yml` in the DSNs and `POSTGRES_PASSWORD=secret` to a stronger DB password.

If you have special characters in your database password, then make sure to URL encode it! It will not work if you have escape characters in your password, such as `<` and `/`.

### Kratos Redirect URLs

In `postgres-nginx/kratos/kratos.yml`'s:

```yaml
serve:
  public:
    base_url: http://localhost:4000/
    cors:
      enabled: true
  admin:
    base_url: http://kratos:4434/
```

and all of the configurations under `selfservice:`, make sure to change all of the `http://localhost:4000/` to the domain you plan on using as the base url to redirect to and use the UI for.

You can also change up the `cors` setings if you need to.

Don't forget to use `https` instead of `http` if you are deploying with NGINX + TLS.

### Email Server Setup

### NGINX HTTPS + Domain Deployment

## Customization
