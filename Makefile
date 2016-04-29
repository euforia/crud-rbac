

.PHONY: clean
clean:
	go clean -i ./...
	rm -f etc/test.bdb

.PHONY: test
test:
	go test -cover ./...
