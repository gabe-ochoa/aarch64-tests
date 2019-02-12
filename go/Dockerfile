FROM golang:1.11

COPY ./go ./

RUN make

CMD [ "./go-aarch64-binary" ]