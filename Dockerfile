FROM golang:1.19

ARG GOPRIVATE=github.com/singhswg

# Change credentials as needed
ARG PG_USER=PG_USER
ARG PG_PASS=PG_PASS
ENV PG_USER=${PG_USER} 
ENV PG_PASS=${PG_PASS}

WORKDIR /app

RUN apt-get update && \
    apt-get install -y \
        git gcc

COPY go.mod go.sum ./

RUN go mod download -x

COPY *.go ./

COPY database ./database

EXPOSE 8080

RUN GOOS=linux GOARCH=amd64 go build -o ./api-db

ENTRYPOINT ["/app/api-db"]