### gorpc

1:
```bash
go mod tidy
```
2: 
```bash
protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. /student/student.proto
```
3:
```bash
go run .
```
