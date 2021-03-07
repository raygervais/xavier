build:
	go build -o bin/xavier-server server/cmd/main.go
	go build -o bin/xavier-client client/cmd/main.go

fmt:
	go fmt ./...

publish: publish-server publish-client

publish-server:
	GOOS=freebsd GOARCH=386 go build -o bin/xavier-server-freebsd-386 server/cmd/main.go
	GOOS=linux   GOARCH=386 go build -o bin/xavier-server-linux-386   server/cmd/main.go
	GOOS=windows GOARCH=386 go build -o bin/xavier-server-windows-386 server/cmd/main.go
	GOOS=darwin  GOARCH=amd64 go build -o bin/xavier-server-macos-64  server/cmd/main.go

publish-client:
	GOOS=freebsd GOARCH=386 go build -o bin/xavier-client-freebsd-386 client/cmd/main.go
	GOOS=linux   GOARCH=386 go build -o bin/xavier-client-linux-386   client/cmd/main.go
	GOOS=windows GOARCH=386 go build -o bin/xavier-client-windows-386 client/cmd/main.go
	GOOS=darwin  GOARCH=amd64 go build -o bin/xavier-client-macos-64  client/cmd/main.go
test:
	go test ./... --vv --cover

race:
	go test ./... --race