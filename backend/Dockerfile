FROM golang:latest

# Set container workdir
WORKDIR /app

# Copy dependency list to container
COPY go.* ./

# Download dependencies
RUN go mod download

# Copy source code to container
COPY . .

# Build the app statically
RUN CGO_ENABLED=0 GOOS=linux go build -o /backend

# Run the app
CMD ["/backend"]
