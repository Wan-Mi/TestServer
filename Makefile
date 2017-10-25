server:
	go build -o "bin/server" ./cmds

run:
	go run ./cmds/main.go

save:
	go dep save ./...