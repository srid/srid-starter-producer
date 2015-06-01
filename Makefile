all:	build run
	@true

build:
	docker build -t srid/demo .

run:
	docker run --rm -it srid/demo
