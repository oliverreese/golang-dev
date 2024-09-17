# syntax=docker/dockerfile:1

FROM golang:1.20-alpine AS build

RUN apk --no-cache add tzdata

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
ADD cmd ./cmd
ADD pkg ./pkg

#RUN go build -o /example ./cmd/example
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /example ./cmd/example

FROM scratch
LABEL org.opencontainers.image.source=https://github.com/oliverreese/golang-dev
WORKDIR /

COPY --from=build /example /example
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Europe/Berlin

CMD [ "/example" ]