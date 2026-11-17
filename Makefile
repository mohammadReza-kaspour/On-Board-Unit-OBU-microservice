.PHONY: obu_sender obu_receiver clean

EXT :=
ifeq ($(OS),Windows_NT)
    EXT := .exe
endif

BIN_DIR := ./bin

obu_sender:
	@mkdir -p $(BIN_DIR)
	@go build -o ./bin/obu_sender$(EXT) ./OBU_Sender
	./bin/obu_sender$(EXT)

obu_receiver:
	@mkdir -p $(BIN_DIR)
	@go build -o ./bin/obu_receiver$(EXT) ./OBU_Receiver
	./bin/obu_receiver$(EXT)

clean:
	rm -rf $(BIN_DIR)
