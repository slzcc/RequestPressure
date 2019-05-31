FROM golang
VOLUME ["/bin"]
COPY ./* $GOPATH/src/
RUN go build -o /rpp $GOPATH/src/rp.go
RUN go build -o /rp $GOPATH/src/rpp.go
ENTRYPOINT /rp
