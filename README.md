# Deploying Fullstack Applications with Ory Kratos' Production Authentication + Social Sign-In

Collection of deployments with a Next.js UI, various backend APIs, various databases, Ory Kratos, and Docker-Compose/K8s.

These examples are intended to be starting points for you to branch off of. The corresponding documentation will help you familiarize yourself with the Ory ecosystem and how it fits into the deployment process!

## Available Workflows

- [Deploy Next.js UI + Echo Go API + Ory Kratos + Postgres with Docker-Compose and NGINX](./docs/next-echo-compose)
- [Kratos Quickstarts + Non-Dockerized Go API + Non-Dockerized Custom Next.js UI](./docs/kratos-quickstart-custom-ui-api)

## Architecture

1. Any requests that have to do with auth will be sent to the Ory Kratos server.
2. Any protected backend ednpoints will be checked by Express/Go API using the Kratos Go SDK before allowing access.
   1. All requests to the backend `localhost:3000` will be proxied as `localhost:4000`
3. Any protected frontend endpoints will check the validity of the session cookie with the Kratos JS SDK.
   1. All requests to the frontend `0.0.0.0:4455` will not be proxied. However, you should update:
      1. `UI_BASE_URL` to match the frontend URL in the API `.env`
      2. The urls in `Ory Console \ Custom UI` to match the `UI_BASE_URL`

All of the SDKs interact with the Kratos server (whether it be the self-hosted Docker image or through the hosted Ory network).

## To-Dos

- [ ] How-To guide for pushing this repo to production
  - I.e. Changing cookie secret, Email SMTP server setup, nginx TLS, postgres secret generation
- [ ] Summary of current deployment process
- [ ] K8s
- [ ] OIDC
- [ ] Practical API dashboard.
- [ ] Practical protected API endpoint
