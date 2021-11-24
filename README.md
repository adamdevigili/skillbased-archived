# skillbased
[![Go Report Card](https://goreportcard.com/badge/github.com/adamdevigili/skillbased.io)](https://goreportcard.com/report/github.com/adamdevigili/skillbased.io)

[Skillbased](http://skillbased.xyz) is a service that is aimed to provide players of casual, "pick-up" type sports and activities to quickly and easily create balanced teams, and save those teams to create elevated levels of competition. Skillbased is designed to be generic,
allowing an unlimited number of custom `Sports` and `Skills` to be added and modified.

## Development

Rename the `.env` file to `.env.local`. The default values should allow the API and postgres to start, however to
enable Auth0 on the frontend, you'll need to supply your own credentials to their corresponding environment variables.

The `DEV_MODE` variable is set to `true` by default, and disables SSL for Postgres.

Build and run the API, frontend, and a postgres database locally with: `docker-compose -f docker-compose.local.yml up --build`

## Tech
### Stack
- Frontend: [NextJS w/ Auth0](https://nextjs.org/)
- API: [Go](https://golang.org/)
- Database: [PostgreSQL](https://www.postgresql.org/)
- Runtime: [Docker](https://www.docker.com/)
- Infrastructure: [DigitalOcean](https://www.digitalocean.com/)

### Other tools
- API designer: [insomnia.rest](https://insomnia.rest/)
- Diagraming: [lucid.app](https://lucid.app/)
