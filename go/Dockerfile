FROM golang:1.11

COPY ./go ./

RUN make build

CMD [ "./go-aarch64-binary" ]