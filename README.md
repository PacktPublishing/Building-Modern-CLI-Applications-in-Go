# audiofile
In Chapter 11 we discuss how to use build tags and testing.  Build tags are setup to separate out the code to be included during the compilation process.  Test files for many of the commands have been added as well.

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

