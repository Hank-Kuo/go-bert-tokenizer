test:
	go test -race -coverpkg=. .

test-bench:
	go test -bench .

run-example:
	go run example/main.go