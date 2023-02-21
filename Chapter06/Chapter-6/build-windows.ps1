Write-Output "installing for windows..." 
Write-Output "installing uppercase..." 
go install .\cmd\uppercase\uppercase.go
Write-Output "installing lettercount..." 
go install .\cmd\lettercount\lettercount.go
Write-Output "installing pages..." 
go install .\cmd\pages\pages.go
Write-Output "installing timeout..." 
go install .\cmd\timeout\timeout.go
Write-Output "installing panic..." 
go install .\cmd\panic\panic.go
Write-Output "installing error..." 
go install .\cmd\error\error.go
Write-Output "installing api..." 
go install .\cmd\api\api.go