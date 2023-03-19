FROM golang:1.20

WORKDIR /go/src

# To run your container withput crash for tests
CMD [ "tail", "-f", "/dev/null" ]
# CMD [ "go", "run", "cmd/app/main.go", "-b", "0.0.0.0" ]
