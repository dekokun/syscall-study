FROM golang:1.6-onbuild
RUN apt-get update
RUN apt-get install strace
ENTRYPOINT ["go"]
CMD ["run", "main.go", "server", ":8000"]
