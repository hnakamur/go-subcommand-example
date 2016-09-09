NAME := go-subcommand-example
VERSION := $(shell git describe --tags)
LDFLAGS := -X 'main.name=$(NAME)' \
           -X 'main.version=$(VERSION)'

all:
	go build -ldflags "$(LDFLAGS)" -o $(NAME)

clean:
	-rm $(NAME)
