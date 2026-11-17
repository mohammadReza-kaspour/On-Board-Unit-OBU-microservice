.PHONY: obu_sender obu_receiver consumer clean

EXT :=
ifeq ($(OS),Windows_NT)
    EXT := .exe
endif

BIN_DIR := ./bin

obu_sender:
	@clear
	@mkdir -p $(BIN_DIR)
	@go build -o ./bin/obu_sender$(EXT) ./OBU_Sender
	./bin/obu_sender$(EXT)

obu_receiver:
	@clear
	@mkdir -p $(BIN_DIR)
	@go build -o ./bin/obu_receiver$(EXT) ./OBU_Receiver
	./bin/obu_receiver$(EXT)

consumer:
	@clear
	@mkdir -p $(BIN_DIR)
	@go build -o ./bin/kafka_consumer$(EXT) ./Kafka_Consumer
	./bin/kafka_consumer$(EXT)

clean:
	rm -rf $(BIN_DIR)
