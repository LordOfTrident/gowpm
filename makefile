BIN = ./bin
APP = $(BIN)/app
CMD = ./cmd/gowpm

GO = go

compile:
	$(GO) build -o $(APP) $(CMD)

$(BIN):
	mkdir -p $(BIN)

run:
	$(GO) run $(CMD)

install:
	cp $(APP) /usr/bin/gowpm

clean:
	rm -r $(BIN)/*

all:
	@echo compile, run, install, clean
