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

# Next.js + Ory Flow

1. Clone next.js + ory example
2. Try it with the ory playground API
3. Then self-host Ory Kratos with Docker
