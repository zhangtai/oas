install:
	mkdir -p ${HOME}/.local/var/oas/{pb_data,log} ${HOME}/.local/bin; \
	touch ${HOME}/.local/var/oas/.env; \
	touch ${HOME}/.local/var/oas/log/oas.{log,err}
	go build -ldflags "-s -w" -o ${HOME}/.local/bin/oas; \
	envsubst < ai.zhangt.oas.plist > ${HOME}/Library/LaunchAgents/ai.zhangt.oas.plist; \
	launchctl enable gui/${UID}/ai.zhangt.oas;

run:
	go run . serve --http localhost:8089

release:
	go build -ldflags "-s -w" -o ${HOME}/.local/bin/oas

restart:
	launchctl kickstart -kp gui/${UID}/ai.zhangt.oas

start:
	launchctl kickstart gui/${UID}/ai.zhangt.oas
