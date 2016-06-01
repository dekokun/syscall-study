FROM golang:1.6
RUN apt-get update && apt-get install strace
RUN mkdir -p /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app
RUN go-wrapper download
RUN go-wrapper install

ENTRYPOINT ["strace"]
CMD ["go", "run", "main.go", "-type=server", "-service=:8000"]
