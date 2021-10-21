FROM golang:latest as build-env

WORKDIR /skillbased

COPY main.go ./
COPY go.mod ./
COPY go.sum ./
COPY . .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

RUN go build -o bin/

FROM alpine:latest
# Go binary dependencies
RUN apk add --no-cache libc6-compat

COPY --from=build-env /skillbased/bin/api /bin/api

EXPOSE 8080

ENTRYPOINT ["/bin/api"]
