build: clean
	@go build

install: clean
	@go install

dist: clean
	@./dist.sh

clean:
	@go clean
	@rm -rf dmsd*.zip

.PHONY: build install dist clean
