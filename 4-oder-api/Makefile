BINPATH = bin/myapp

.PHONY: build
build:
	go build -o $(BINPATH) cmd/app/main.go

.PHONY: run
run:	
	@if [ ! -e $(BINPATH) ]; then go build -o $(BINPATH) cmd/app/main.go; fi; $(BINPATH)

.PHONY: rebuild
rebuild:	
	@if [ -e $(BINPATH) ]; then rm -r $(BINPATH); fi; go build -o $(BINPATH) cmd/app/main.go; $(BINPATH)

