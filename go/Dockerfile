FROM golang:1.11 as build

COPY ./go ./

RUN make build

FROM alpine

COPY --from=build /go/go-aarch64-binary /go-aarch64-binary

ENTRYPOINT [ "/go-aarch64-binary" ]