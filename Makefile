test:
	go test ./...

bench:
	go test -bench=. ./...

update:
	leetcode update
	leetcode readme
	leetcode helper
	leetcode tag
	leetcode description

build:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
