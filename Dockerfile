FROM golang:1.16.7-alpine3.13 as build
RUN mkdir -p /app
WORKDIR /app
COPY . .

ENV CGO_ENABLED=0

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh gcc musl-dev && \
    go mod download && \
    go build -o main cmd/api/main.go


FROM alpine:3.13 as prod
WORKDIR /

RUN mkdir -p /home/appuser/app
COPY --from=build /app/main /home/appuser/app/main

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
RUN chown -R appuser:appgroup /home/appuser
USER appuser

ENTRYPOINT /home/appuser/app/main