FROM golang:latest as build-env
# All these steps will be cached

WORKDIR /skillbased

COPY main.go ./
COPY go.mod ./
COPY go.sum ./
COPY ./api ./api

COPY ./frontend/dist ./frontend/dist

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

RUN go build -o bin/

FROM alpine:latest
# Go binary dependencies
RUN apk add --no-cache libc6-compat

COPY --from=build-env /skillbased/bin/skillbased /bin/skillbased

ENTRYPOINT ["/bin/skillbased"]