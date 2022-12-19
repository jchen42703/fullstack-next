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

## To-Dos

- [ ] Practical Go Endpoint setup with Postgres (Sequelize)
  - Don't redirect, just return 401
  - Echo, JSON Middleware
- [ ] Move next-js auth api handlers to Go
  - https://github.com/atreya2011/go-kratos-test
  - https://github.com/ory/hydra/discussions/2873
- [ ] Email SMTP Server Setup
- [ ] Practical UI Endpoint protection with protected dashboard.
