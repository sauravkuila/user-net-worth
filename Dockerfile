#Build stage
FROM golang:1.20-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o ./app/app ./app

#Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app /
EXPOSE 8080
RUN adduser --disabled-password --gecos "" --uid 1024 alpha
RUN chown alpha:alpha /app /var/log -R
USER alpha
CMD ./app