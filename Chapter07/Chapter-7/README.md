# Chapter-7

## Preparing to run the code on MacOS
Before you run the main code, build the sleep command application:
`make install`

## Running the code on MacOS
Once the `$GOPATH/bin` folder contains the executable, you're ready to run the code:
`make run`

## Preparing to run the code on Windows
Before you run the main code, build the sleep command application:
`.\build-windows.ps1`

## Running the code on Windows
Once the `$GOPATH/bin` folder contains the executable, you're ready to run the code:
`.\run-windows.ps1`

## Preparing to run the code on Linux
Before you run the main code, build the sleep command application:
`chmod +x build-linux.sh`
`./build-linux.sh`

## Running the code on Linux
Once the `$GOPATH/bin` folder contains the executable, you're ready to run the code:
`./run-linux.sh`


Since the main function contains calls to all functions, feel free to comment out and run each separately to view less output at a time.