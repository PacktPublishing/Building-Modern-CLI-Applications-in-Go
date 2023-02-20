# Chapter-13

In Chapter 13, we discuss how to use Docker containers to test and distribute command line applications.  Here we go over a very basic example of how to create a Docker image and then run a Docker container to execute a simple Go application that prints "Hello, World!" when the --hello flag is receieved.

## Build image
docker build --tag hello-world:latest .

## Run container
docker run hello-world:latest