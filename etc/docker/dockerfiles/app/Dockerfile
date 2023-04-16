# Importing oficial Golang image
FROM golang:1.20

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /pelando

EXPOSE 8080

# Run
CMD ["/pelando"]
