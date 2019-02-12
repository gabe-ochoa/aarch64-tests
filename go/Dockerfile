FROM golang:1.11 as build

COPY ./go ./

RUN GOARCH=arm64 GOOS=linux CGO_ENABLED=0 go build -o go-aarch64-binary -tags netgo -a -v

CMD [ "/go/go-aarch64-binary" ]