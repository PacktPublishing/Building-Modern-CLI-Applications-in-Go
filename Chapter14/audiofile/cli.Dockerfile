FROM golang:1.19

# Set the working directory
WORKDIR /audiofile

# Copy the source code
COPY . .

# Download the dependencies
RUN go mod download

# Execute `go test -v ./cmd -tags int pro` when the container is running
CMD ["go", "test", "-v", "./cmd", "-tags", "int pro"]