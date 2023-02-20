# audiofile
In Chapter 6, we discuss about handling timeouts an errors.  In this next revision of the command line interface, we add a function to check the response from the client and add a timeout value of fifteen seconds to the client.

## To generate the audiofile command line interface:
make build

## To start the audiofile API (required for the CLI to run):
The API must be started and running before the CLI.  Start the APi in a separate terminal.  Within the working directory:
./bin/audiofile api

## To call the audiofile command line interface:
./bin/audiofile

