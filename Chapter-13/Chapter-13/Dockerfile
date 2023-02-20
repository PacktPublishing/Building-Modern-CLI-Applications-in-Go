FROM golang:1.19

# Copy all the files
COPY . .

# Build the hello world application
RUN go build main.go

# Run `./main --hello`
CMD ["./main", "--hello"]