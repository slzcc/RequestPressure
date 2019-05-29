FROM golang
VOLUME ["/bin"]
COPY ./* $GOPATH/src/
RUN go build -o /rq $GOPATH/src/rp.go
RUN go build -o /rq $GOPATH/src/rpp.go
ENTRYPOINT /rp