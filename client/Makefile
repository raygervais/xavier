build:
	go build -o bin/xavier cmd/xavier/main.go

run:
	go run cmd/xavier/main.go

fmt:
	go fmt ./...

publish:
	GOOS=freebsd GOARCH=386 go build -o bin/xavier-freebsd-386 cmd/gophernotes/main.go
	GOOS=linux   GOARCH=386 go build -o bin/xavier-linux-386   cmd/gophernotes/main.go
	GOOS=windows GOARCH=386 go build -o bin/xavier-windows-386 cmd/gophernotes/main.go
	GOOS=darwin  GOARCH=amd64 go build -o bin/xavier-macos-64  cmd/gophernotes/main.go

test:
	go test ./... --vv --cover

race:
	go test ./... --race