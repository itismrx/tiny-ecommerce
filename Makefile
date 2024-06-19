build:
	@~/go/bin/air  

run: build
	@./bin/tiny-ecommerce

test:
	@go test -v ./..@