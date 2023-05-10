build:
	@go build -o ./bin/bankserver

run:
	@./bin/bankserver

clean:
	@rm -rf ./bin