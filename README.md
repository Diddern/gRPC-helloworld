# gRPC-simpleGCDService

A small service to test gRPC

## Installing

```
git clone https://github.com/Diddern/gRPC-simpleGCDService && cd gRPC-simpleGCDService/
```

## Running

Create certificates:
```
openssl req -x509 -newkey rsa:4096 -keyout certs/server-key.pem -out certs/server-cert.pem -days 365 -nodes -subj '/CN=localhost'
```

Running the server:  
This will start the server and listen for gRPC-requests on port 3000.
```
go run gcd/main.go
```

Calculate the GCD of 294 and 462:

```
go run client/main.go 294 462
```

## If you need to rebuild the protobuf

Navigate to the pb/ directory and run the following command:

```
protoc -I . --go_out=plugins=grpc:. ./*.proto
```

Compilation should produce gcd.pb.go file.
