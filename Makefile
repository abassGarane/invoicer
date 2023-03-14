.PHONY: build, run, sqlc, db_up, db_down, clean

build:
	@ CGO_ENABLE=0 go build -o bin/learning -ldflags "-s -w" main.go

run: build
	@ ./bin/learning

sqlc:
	@echo "Generating stubs..."
	@sqlc generate


db_up:
	sudo docker-compose up -d

db_down:
	sudo docker-compose down

clean:
	rm -rf ./bin
