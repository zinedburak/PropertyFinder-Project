FROM golang:1.17

WORKDIR /PropertyFinder-Project
COPY ./ /PropertyFinder-Project

RUN go get ./...


EXPOSE 8000

CMD [ "go", "run", "cmd/main.go" ]