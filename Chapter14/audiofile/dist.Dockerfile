FROM golang:1.19

# Set the working directory
WORKDIR /audiofile

# Copy the source code
COPY . .

# Download the dependencies
RUN go mod download

# Expose port 8000
EXPOSE 8000

# Build the audiofile application with the pro tag so all features are available
RUN go build -tags "pro" -o audiofile main.go

# Start the audiofile API
ENTRYPOINT ["./audiofile"]