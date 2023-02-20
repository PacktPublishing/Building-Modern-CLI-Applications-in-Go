# audiofile
In Chapter 12, Cross Compilation Across Different Platforms , we discuss how to cross compile for the three major platforms: darwin, linux, and windows.  We also write a couple scripts to expedite compilation for all possible platforms and architectures.  Visit the Makefile to see all the new commands!

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

