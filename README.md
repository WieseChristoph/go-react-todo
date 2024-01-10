# Go-React Todo

Simple todo app with Discord authentication to try out Go as a backend.

## Production (Docker)

1. Copy `.env.example` to `.env`
2. Set all variables in `.env`
3. Start the docker compose
   ```bash
   docker compose up [-d]
   ```

> [!NOTE]
> The PostgreSQL data is stored in `postgres-data/`

## Development (Docker)

1. Go to `_DEV/`
2. Copy `.env.example` to `.env`
3. Set all variables in `.env`
4. Start the docker compose
   ```bash
   docker compose up [-d]
   ```
5. Access the site on default port 8080

> [!NOTE]
> The backend and frontend will reload automatically on file changes

## Environment variables

| Name                  | Description                                                                | Example             |
| --------------------- | -------------------------------------------------------------------------- | ------------------- |
| APP_URL               | The URL where the app will be hosted                                       | https://example.com |
| DISCORD_CLIENT_ID     | Discord OAuth2 client ID (https://discord.com/developers/applications)     |                     |
| DISCORD_CLIENT_SECRET | Discord OAuth2 client secret (https://discord.com/developers/applications) |                     |