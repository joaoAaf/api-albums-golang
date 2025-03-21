FROM golang:1.21.0-alpine AS build
WORKDIR /app
COPY api/go.mod api/go.sum ./
RUN go mod download && go mod verify
COPY api .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./api
FROM gcr.io/distroless/static-debian12 AS realease
COPY --from=build /app/api /usr/local/bin/api
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT [ "api" ]