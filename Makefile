run_server:
	go run cmd/server/main.go
run_client:
	go run cmd/client/main.go
run_client_gui:
	python3 cmd/client/python-gui/main.py

pre_run:
	cp config/server.toml.bak config/server.toml
	cp config/client.toml.bak config/client.toml
	sed -i "s/Secret = \"\"/Secret = \"$(python3 scripts/get_secret.py)\"/g" config/server.toml
	sed -i 's/LogFile = "\.\/logs\/server\.log"/LogFile =""/g' config/server.toml

build_win64:
	GOOS=windows go build -o build/win_x86/golearn-server.exe cmd/server/main.go
	GOOS=windows go build -o build/win_x86/golearn-client.exe cmd/client/main.go
build_lin64:
	GOOS=linux go build -o build/lin_amd64/golearn-server cmd/server/main.go
	GOOS=linux go build -o build/lin_amd64/golearn-client cmd/client/main.go
build_mac64:
	GOOS=darwin go build -o build/mac_amd64/golearn-server cmd/server/main.go
	GOOS=darwin go build -o build/mac_amd64/golearn-client cmd/client/main.go

set_version:
	python3 scripts/set_version.py