# syntax=docker/dockerfile:1

FROM golang:1.21 AS build

# Set destination for COPY
WORKDIR /src

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY ./cmd/main.go ./cmd/main.go

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /backend ./cmd/main.go

# Use light weight image for prod later 
FROM golang:1.21

WORKDIR /

COPY --from=build /backend /backend

EXPOSE 8000

# Run
CMD ["/backend"]
