# audiofile
In Chapter 4, we discuss abou the frameworks that can speed up CLI development.  Within this Chapter-4 branch, we convert the older version of the command line interface to utilize the Cobra package.

## To generate the audiofile command line interface:
make build

## To start the audiofile API (required for the CLI to run):
The API must be started and running before the CLI.  Start the APi in a separate terminal.  Within the working directory:
./bin/audiofile api

## To call the audiofile command line interface:
./bin/audiofile

