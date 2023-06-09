NAME=dop-exchange
VERSION=0.0.1

.PHONY: build
## build: Compile the packages.
build:
	@go build -o $(NAME)

.PHONY: run
## run: Build and Run in development mode.
run: build
	@./$(NAME) -e development

.PHONY: clean
## clean: Clean project and previous builds.
clean:
	@rm -f $(NAME)

.PHONY: deps
## deps: Download modules
deps:
	@go mod download

.PHONY: gen-doc
## gen-doc: Generate Swagger documentation
gen-doc:
	@swag init

## docker-build: Builds and tag docker image
docker-build:
	@docker build -t luisjmarrero/dop-exchange . 


## docker-run: Run docker container
docker-run:
	@docker run -p 8080:8080 luisjmarrero/dop-exchange

.PHONY: help
all: help
# help: show this help message
help: Makefile
	@echo
	@echo " Choose a command to run in "$(APP_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo