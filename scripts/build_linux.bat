cd ..
del todo_backend
set GOOS=linux
set GOARCH=amd64
go build -o todo_backend main.go
