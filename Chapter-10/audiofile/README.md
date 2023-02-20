# audiofile
In Chapter 10 we discuss how to create a terminal dashboard.  Here we add a new command, player, which launches a audiofile player dashboard in the users terminal.

## To generate the audiofile CLI documentation on MacOS:
make manpages

## To generate the audiofile command line interface on MacOS:
make build-darwin

## [UNIX] To start the audiofile API (required for the CLI to run):
The API must be started and running before the CLI.  Start the APi in a separate terminal.  Within the working directory:
./bin/audiofile api

## [UNIX] To call the audiofile command line interface:
./bin/audiofile

