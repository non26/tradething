
build:
	docker build . --tag debug-image --file ../Dockerfile.debug

run:
	docker run --publish 80:80 --publish 4000:4000 --name debug-server debug-image