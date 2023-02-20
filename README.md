# audiofile
In Chapter 3, we discuss a command line interface built from scratch which handles generating metadata from uploaded audio files, local flat file storage and retrieval of audio metadata.  This CLI is just an example and for reference to the chapter.  It was created on MacOS and other operating systems have not yet been tested at this time.

## To generate the audiofile command line interface:
go build -o audiofile-cli cmd/cli/main.go

## To generate the audiofile API:
go build -o audiofile-api cmd/api/main.go

## Within the root of the audiofile folder, to start the API:
./audiofile-api

### NOTE
To change the default port, 8000, pass in the new port value with the `-p` flag.

## To call the audiofile command line interface:
./audiofile-cli
