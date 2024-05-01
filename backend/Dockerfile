FROM golang:latest

# Set container workdir
WORKDIR /app

# Update and install external dependencies
RUN apt-get update && apt-get install -y libimage-exiftool-perl

# Copy dependency list to container
COPY go.* ./

# Download dependencies
RUN go mod download && go mod verify

# Copy source code to container
COPY . .

# Build the app statically
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /app.out

# Run the app
CMD ["/app.out"]
