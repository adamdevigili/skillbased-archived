# skillbased.io
[![Go Report Card](https://goreportcard.com/badge/github.com/adamdevigili/skillbased.io)](https://goreportcard.com/report/github.com/adamdevigili/skillbased.io)

[skillbased](http://skillbased.xyz) is a service that is aimed to provide players of casual, "pick-up" type sports to quickly and easily create balanced teams for their activity, and save those teams to create elevated levels of competition.

## Development

Rename the `.env` file to `.env.local`. The default values should allow the API and postgres to start, however to
enable Auth0 on the frontend, you'll need to supply your own credentials to their corresponding environment variables

Build and run the API, frontend, and a postgres database locally with: `docker-compose -f docker-compose.local.yml up --build`

## Tech
### Stack
- Frontend: [ReactJS](https://reactjs.org/)
- Backend: [Go](https://golang.org/) ([echo](https://echo.labstack.com/))
- Database: [PostgreSQL](https://www.postgresql.org/)
- Runtime: [Docker](https://www.docker.com/)
- Infrastructure: [DigitalOcean](https://www.digitalocean.com/)


### Other tools
- API designer: [insomnia.rest](https://insomnia.rest/)
