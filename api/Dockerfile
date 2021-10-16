FROM golang:latest as build-env
# All these steps will be cached
WORKDIR /api

COPY . .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

WORKDIR cmd/api

RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo

FROM alpine:latest
COPY --from=build-env /api/cmd/api /
ENTRYPOINT ["./api"]


# --- WORKING ---
#FROM golang:latest
#
#ENV GO111MODULE=on
#WORKDIR /app
#COPY ./go.mod .
#RUN go mod download
#COPY . .
#CMD ["go", "run", "cmd/api/main.go"]









#FROM golang:latest as build-env
## All these steps will be cached
#RUN mkdir /app
#WORKDIR /app
#COPY go.mod .
#COPY go.sum .
#
## Get dependancies - will also be cached if we won't change mod/sum
#RUN go mod download
## COPY the source code as the last step
#COPY . .
#
## Build the binary
#RUN cd cmd/api CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app
#
##FROM alpine
##COPY --from=build-env /go/bin/app /go/bin/app
#ENTRYPOINT ["/go/bin/app"]