
NOW = $(shell date  '+%Y-%m-%d %H:%M:%S')

.PHONY: check dist build run all check

all: build

check: test all build clean fmt todo legacy


build:
	go build  -ldflags "-X  'github.com/zhangshanwen/the_one/api/v1.version=$(NOW)' "  -o bin/shard  cmd/api.go

run:build
	./bin/shard


clean:
	find . -name "*.DS_Store" -type f -delete
	rm -rf bin

test:
	go test -cover -race ./...


fmt:
	go fmt  ./...

todo:
	grep -rnw "TODO" internal

# Legacy code should be removed by the time of release
legacy:
	grep -rnw "\(LEGACY\|Deprecated\)" internal
