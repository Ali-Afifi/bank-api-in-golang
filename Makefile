build:
	@go build -o ./bin/bankserver

run: build
	@./bin/bankserver

clean:
	@rm -rf ./bin

db:
	docker compose -f database-compose.yml up

compose:
	docker compose up --build --force-recreate                                                                                    	