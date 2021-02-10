# build container
FROM golang:1.15-alpine as builder
WORKDIR /app
COPY . .
RUN go install
RUN go build server.go

# app container
FROM alpine
WORKDIR /app
COPY --from=builder /app/server .
CMD [ "/app/server" ]
