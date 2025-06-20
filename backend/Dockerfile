# syntax=docker/dockerfile:1

FROM golang:1.24 AS build

# Set destination for COPY
WORKDIR /src

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY . .

# Use env until find way to mount /run/secrets
# COPY .env .env


# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /backend ./cmd/main.go

# Use light weight image for prod later 
FROM golang:1.24

WORKDIR /

COPY --from=build /backend /backend
COPY --from=build /src/.env .

EXPOSE 8000
ARG APP_ENV
# Run
CMD ["/backend" "APP_ENV=docker-test"]
