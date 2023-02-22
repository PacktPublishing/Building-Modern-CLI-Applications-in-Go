# Build stage
FROM golang:1.19 AS build

# Set the working directory
WORKDIR /audiofile

# Copy the source code
COPY . .

# Download the dependencies
RUN go mod download

# Build the audiofile application with the pro tag so all features are available
RUN go build -tags "pro" -o audiofile main.go

# Final stage
FROM alpine:latest

# Copy the compiled binary from the build stage to the final stage
COPY --from=build /audiofile/audiofile .

# Expose port 8000
EXPOSE 8000

# Start the audiofile API
ENTRYPOINT ["./audiofile"]
