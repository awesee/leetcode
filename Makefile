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
