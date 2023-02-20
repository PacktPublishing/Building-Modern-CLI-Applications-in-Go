# audiofile
In Chapter 7, we discuss about how to build for multiple platforms.  You'll see the addition of files that are built specific to certain plaforms.  However, cross compilation will be handled in later chapters.  This build only has a make build command for darwin only.

## To generate the audiofile command line interface on MacOS:
make build-darwin

## [UNIX] To start the audiofile API (required for the CLI to run):
The API must be started and running before the CLI.  Start the APi in a separate terminal.  Within the working directory:
./bin/audiofile api

## [UNIX] To call the audiofile command line interface:
./bin/audiofile

