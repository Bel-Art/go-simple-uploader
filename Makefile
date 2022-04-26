all: docker run

docker:
	docker build -t n4n5/go-simple-uploader:1.0.0 .

run:
	docker run -p 7000:8080 go-simple-uploader

push:
	# use docker login
	docker push n4n5/go-simple-uploader:1.0.0