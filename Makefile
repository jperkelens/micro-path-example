default: run

_build-micro:
	cd $$GOPATH/src/github.com/micro/micro && \
		echo 'package main \n\nimport _ "github.com/micro/go-plugins/client/grpc"\nimport _ "github.com/micro/go-plugins/server/grpc"\nimport _ "github.com/micro/go-plugins/registry/kubernetes"' > plugins.go && \
		go build -i -o micro main.go plugins.go
		mv $$GOPATH/src/github.com/micro/micro/micro $$GOPATH/bin

build: clean lint
	cd server && CGO_ENABLED=0 go build -a -installsuffix cgo -o main main.go

build-linux: clean lint
	cd server && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go

clean:
	rm server/main || :

config-local:
	which glide || brew install glide
	which protoc || brew install protobuf
	which node || brew install node
	which yarn || npm install -g yarn
	go get -u \
		github.com/micro/micro \
		github.com/micro/go-plugins/client/grpc \
		github.com/micro/go-plugins/server/grpc \
		github.com/micro/go-plugins/registry/kubernetes \
		github.com/micro/protobuf/{proto,protoc-gen-go} \
		gopkg.in/alecthomas/gometalinter.v2 
	gometalinter.v2 --install
	$(MAKE) _build-micro
	glide install
	cd client && yarn install

lint:
	gometalinter.v2 ./...

protoc:
	protoc --go_out=plugins=micro:. proto/*.proto

run:
	./server/main

