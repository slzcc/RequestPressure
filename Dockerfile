FROM golang
VOLUME ["/bin"]
COPY ./* $GOPATH/src/
RUN go build -o /rq $GOPATH/src/rq.go
ENTRYPOINT /rq