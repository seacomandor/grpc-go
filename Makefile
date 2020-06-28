all: vet test testrace

build: deps
	go build github.com/seacomandor/grpc-go/...

clean:
	go clean -i github.com/seacomandor/grpc-go/...

deps:
	go get -d -v github.com/seacomandor/grpc-go/...

proto:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	go generate github.com/seacomandor/grpc-go/...

test: testdeps
	go test -cpu 1,4 -timeout 7m github.com/seacomandor/grpc-go/...

testsubmodule: testdeps
	cd security/advancedtls && go test -cpu 1,4 -timeout 7m github.com/seacomandor/grpc-go/security/advancedtls/...

testappengine: testappenginedeps
	goapp test -cpu 1,4 -timeout 7m github.com/seacomandor/grpc-go/...

testappenginedeps:
	goapp get -d -v -t -tags 'appengine appenginevm' github.com/seacomandor/grpc-go/...

testdeps:
	go get -d -v -t github.com/seacomandor/grpc-go/...

testrace: testdeps
	go test -race -cpu 1,4 -timeout 7m github.com/seacomandor/grpc-go/...

updatedeps:
	go get -d -v -u -f github.com/seacomandor/grpc-go/...

updatetestdeps:
	go get -d -v -t -u -f github.com/seacomandor/grpc-go/...

vet: vetdeps
	./vet.sh

vetdeps:
	./vet.sh -install

.PHONY: \
	all \
	build \
	clean \
	deps \
	proto \
	test \
	testappengine \
	testappenginedeps \
	testdeps \
	testrace \
	updatedeps \
	updatetestdeps \
	vet \
	vetdeps
