buf-install:
	curl -sSL https://github.com/bufbuild/buf/releases/latest/download/buf-Darwin-arm64 -o buf # install for mac m+
	chmod +x buf
	sudo mv buf /usr/local/bin/buf
	buf --version

servergen:
	buf generate proto --config bufconfig/server/buf.yaml --template bufconfig/server/buf.gen.yaml