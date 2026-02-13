buf-install:
	curl -sSL https://github.com/bufbuild/buf/releases/latest/download/buf-Darwin-arm64 -o buf # install for mac m+
	chmod +x buf
	sudo mv buf /usr/local/bin/buf
	buf --version

server-gen:
	buf generate --config bufconfig/server/buf.yaml --template bufconfig/server/buf.gen.yaml

client-gen:
	buf generate --template bufconfig/client/buf.gen.yaml

run-server:
	go run cmd/server/server.go

run-client:
	go run cmd/client/client.go