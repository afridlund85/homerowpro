all: build

init: 
	cp /usr/lib/go/lib/wasm/wasm_exec.js ./web/public/

clean: 
	rm -rf dist
	mkdir dist

build: build-wasm build-webserver

build-wasm: clean
	cd cmd/wasm && GOOS=js GOARCH=wasm go build -o ../../dist/app.wasm

build-webserver: clean
	cd cmd/webserver && go build -o ../../dist/webserver

