FROM golang:alpine

RUN apk update
RUN apk upgrade
RUN apk add bash
RUN apk add git

# Set necessary environmet variables needed for our image
ENV GO111MODULE=auto \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Generate swagger docs
RUN go install github.com/swaggo/swag/cmd/swag@latest
#COPY --from=itinance/swag /root/swag /usr/local/bin
RUN swag init --parseDependency --parseInternal --parseDepth 1

# Build the application
RUN go build -o main main.go

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist
COPY config/config.json config/config.json
# Copy binary from build to main folder
RUN cp /build/main .

# Command to run when starting the container
CMD ["/dist/main"]