.PHONY: build
build:
	(if exist build rd -s -q build) && mkdir build && go build -o build/generator.exe -v cmd/generator/main.go

.PHONY: run
run:
	go run cmd/generator/main.go

serve:
	build/generator.exe

.DEFAULT_GOAL := build