FROM golang
COPY . /app
RUN cd /app ; go build -o go-simple-uploader .
CMD cd /app; ./go-simple-uploader