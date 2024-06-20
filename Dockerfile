# Build.
FROM golang:1.22 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint ./cmd/http

# Deploy.
FROM gcr.io/distroless/static-debian11 AS release-stage

WORKDIR /

COPY --from=build-stage /entrypoint /entrypoint
COPY --from=build-stage /app/assets /assets

EXPOSE 8080

USER nonroot:nonroot

ENV DB_DNS="user=user password=pw dbname=postgres port=5432 host=localhost"
ENTRYPOINT ["/entrypoint"]
