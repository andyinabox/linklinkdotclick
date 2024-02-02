# syntax=docker/dockerfile:1
# FROM golang:1.18 AS builder
FROM golang:1.18-alpine

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app main.go

# go build -o dist/linkydink main.go
# RUN CGO_ENABLED=1 GOOS=linux go build -v -o /usr/local/bin/app -a -ldflags '-linkmode external -extldflags "-static"' main.go

# FROM scratch
# COPY --from=builder /usr/local/bin/app /usr/local/bin/app
EXPOSE 8080

# ENTRYPOINT ["/usr/local/bin/app"]
CMD ["app", "--dbfile=/db/data.db", "--mode=release"]