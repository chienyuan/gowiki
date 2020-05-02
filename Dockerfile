FROM golang:1.14 as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
#RUN go get -d -v github.com/nsf/termbox-go
#RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o gowiki  .
RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o gowiki *.go

FROM scratch
COPY --from=builder /build/gowiki /app/
COPY --from=builder /build/*.html /app/
EXPOsE 8080
WORKDIR /app
CMD ["./gowiki"]
