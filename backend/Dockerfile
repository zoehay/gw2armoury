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
RUN go build -o backend ./cmd/main.go

RUN chmod +rx backend

FROM golang:1.21

WORKDIR /src

COPY --from=build /src/backend backend
RUN ls -la
EXPOSE 8000

# Run
CMD ["./backend"]
