# audiofile
In Chapter 13, Using Containers for Distribution, we start using Docker containers to handle integration tests, and to distribute our application via Docker Hub. 

## To run the integration tests with Docker Compose:
`docker-compose up`

## To run the api as a container:
`docker build -f api.Dockerfile -t audiofile:api .`
`docker run -rm -p 8000:8000 audiofile:api`

## To run the cli as a container:
`docker build -f cli.Dockerfile -t audiofile:cli .`
`docker run -rm --network host audiofile:cli `

## To run the cli as an executable:
`docker build -f dist.Dockerfile -t audiofile:dist .`
`docker run --rm --network host -ti audiofile:dist help`


## To run the multi-stage Docker build as an executable:
`docker build -f dist-multistage.Dockerfile -t audiofile:dist-multistage .`
`docker run --rm --network host -ti audiofile:dist-multistage help`


## To run the cli as an executable from DockerHub:
`docker run --rm --network host -ti marianmontagnino/audiofile:latest help`

## To run tests:
make test

## To run tests in verbose mode:
make test-verbose

## To generate the audiofile CLI documentation on MacOS:
make manpages

## To generate the audiofile command line interface on MacOS:
make build-darwin

## [UNIX] To start the audiofile API (required for the CLI to run):
The API must be started and running before the CLI.  Start the APi in a separate terminal.  Within the working directory:
./bin/audiofile api

## [UNIX] To call the audiofile command line interface:
./bin/audiofile

