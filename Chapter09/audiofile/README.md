# audiofile
In Chapter 9 we discuss how to integrate empathy into our documentation.  Errors are decorated and reworded before returned.  Also, logging with the Zap logger has been added and the verbose flag shows this debug output when used.  Finally, a new documentation page has been added in conjunction with a make command to generate man pages for the audiofile CLI.

## To generate the audiofile CLI documentation on MacOS:
make manpages

## To generate the audiofile command line interface on MacOS:
make build-darwin

## [UNIX] To start the audiofile API (required for the CLI to run):
The API must be started and running before the CLI.  Start the APi in a separate terminal.  Within the working directory:
./bin/audiofile api

## [UNIX] To call the audiofile command line interface:
./bin/audiofile

