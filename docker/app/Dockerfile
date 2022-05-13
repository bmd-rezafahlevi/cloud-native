# Build environment
# -----------------
FROM golang:1.13-alpine as build-env
WORKDIR /cloud-native

RUN apk update && apk add --no-cache gcc musl-dev git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags ' -w -s' -a -o ./bin/app ./cmd/app \
    && go build -ldflags '-w -s' -a -o ./bin/migrate ./cmd/migrate

# Deployment environment
# ----------------------
FROM alpine
RUN apk update && apk add --no-cache bash mysql-client

COPY --from=build-env /cloud-native/bin/app /cloud-native/
COPY --from=build-env /cloud-native/bin/migrate /cloud-native
COPY --from=build-env /cloud-native/migrations /cloud-native/migrations

COPY --from=build-env /cloud-native/docker/app/bin /usr/local/bin/cloud-native/
RUN chmod +x /usr/local/bin/cloud-native/*

EXPOSE 2405
CMD ["/cloud-native/app"]