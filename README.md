# Deploying Fullstack Applications with Ory Kratos' Production Authentication

Collection of deployments with a Next.js UI, various backend APIs, various databases, Ory Kratos, and Docker Compose/K8s.

These examples are intended to be starting points for you to branch off of. The corresponding documentation will help you familiarize yourself with the Ory ecosystem and how it fits into the deployment process!

## Available Workflows

- [Deploy Next.js UI + Echo Go API + Ory Kratos + Postgres with Docker-Compose and NGINX](./docs/next-echo-compose)
- [Kratos Quickstarts + Non-Dockerized Go API + Non-Dockerized Custom Next.js UI](./docs/kratos-quickstart-custom-ui-api)

## Available Components

- [Docker Compose Deployments](./deployment/compose/)
- [Custom Next.js UI + Ory Kratos](./ui)
- [Echo Go API + Ory Kratos](./echo-api/)
  - [Endpoint Protection Middleware with Ory Kratos](./echo-api/middleware/auth.go)
  - [Protected Endpoint Example](./echo-api/controllers/baseRoutes/router.go)
  - [Non-Protected Endpoint Example](./echo-api/controllers/baseRoutes/router.go)

## To-Dos

- [ ] How-To guide for pushing this repo to production
  - Changing cookie secret, Email SMTP server setup, setting up NGINX TLS, postgres secret generation
- [ ] Summary of current deployment process
- [ ] K8s
- [ ] OIDC
- [ ] Practical API dashboard.
- [ ] Practical protected API endpoint
