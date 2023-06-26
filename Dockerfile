# Start from a base Go image
FROM golang:1.19-alpine

# Set the working directory inside the container , Copy the necessary files and directories into the container & set the working directory to /app
RUN mkdir /app

ADD . /app

WORKDIR /app

# Build the Go application
RUN go build -o main cmd/main.go

# Set the command to run your server when the container starts
CMD ["/app/main"]