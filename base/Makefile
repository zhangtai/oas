run:
	go run . serve --http localhost:8089

release:
	go build -ldflags "-s -w" -o ${HOME}/.local/bin/oas

stop:
	launchctl stop ai.zhangt.oas

start:
	launchctl start ai.zhangt.oas
