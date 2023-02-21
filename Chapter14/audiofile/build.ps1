# Generate darwin builds
$darwin_archs="amd64","arm64"
foreach ($darwin_arch in $darwin_archs)
{
    Write-Output "building for darwin/$($darwin_arch) free version..."  
    $env:GOOS="darwin";$env:GOARCH=$darwin_arch; go build -tags free -o .\builds\free\darwin\$darwin_arch\audiofile main.go
    Write-Output "building for darwin/$($darwin_arch) pro version..."  
    $env:GOOS="darwin";$env:GOARCH=$darwin_arch; go build -tags pro -o .\builds\pro\darwin\$darwin_arch\audiofile main.go
    Write-Output "building for darwin/$($darwin_arch) profile version..."  
    $env:GOOS="darwin";$env:GOARCH=$darwin_arch; go build -tags profile -o .\builds\profile\darwin\$darwin_arch\audiofile main.go
}

# Generate linux builds
$linux_archs="386","amd64","arm","arm64","loong64","mips","mips64","mips64le","mipsle","ppc64","ppc64le","riscv64","s390x"
foreach ($linux_arch in $linux_archs)
{
    Write-Output "building for linux/$($linux_arch) free version..."  
    $env:GOOS="linux";$env:GOARCH=$linux_arch; go build -tags free -o .\builds\free\linux\$linux_arch\audiofile main.go
    Write-Output "building for linux/$($linux_arch) pro version..."  
    $env:GOOS="linux";$env:GOARCH=$linux_arch; go build -tags pro -o .\builds\pro\linux\$linux_arch\audiofile main.go
    Write-Output "building for linux/$($linux_arch) profile version..."  
    $env:GOOS="linux";$env:GOARCH=$linux_arch; go build -tags profile -o .\builds\profile\linux\$linux_arch\audiofile main.go
}

# Generate windows builds
$windows_archs="386","amd64","arm","arm64"
foreach ($windows_arch in $windows_archs)
{
    Write-Output "building for windows/$($windows_arch) free version..." 
    $env:GOOS="windows";$env:GOARCH=$windows_arch; go build -tags free -o .\builds\free\windows\$windows_arch\audiofile.exe main.go
    Write-Output "building for windows/$($windows_arch) pro version..." 
    $env:GOOS="windows";$env:GOARCH=$windows_arch; go build -tags pro -o .\builds\pro\windows\$windows_arch\audiofile.exe main.go
    Write-Output "building for windows/$($windows_arch) profile version..." 
    $env:GOOS="windows";$env:GOARCH=$windows_arch; go build -tags profile -o .\builds\profile\windows\$windows_arch\audiofile.exe main.go
}