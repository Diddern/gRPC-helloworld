# gRPC-simpleGCDService

A small service to test gRPC

## Installing

``` 
go get https://github.com/Diddern/gRPC-simpleGCDService 
``` 

##Running

Start the server:

``` go run gcd/main.go ```

Calculate the GCD of 294 and 462:

``` go run client/main.go 294 462 ```