FROM golang:1.11 as build

COPY ./go ./

# Get QEMU handlers
RUN wget -N https://github.com/multiarch/qemu-user-static/releases/download/v2.9.1-1/x86_64_qemu-aarch64-static.tar.gz -O - | tar --totals -xz -C /usr/bin

RUN GOARCH=arm64 GOOS=linux CGO_ENABLED=0 go build -o go-aarch64-binary -tags netgo -a -v

CMD [ "/go/go-aarch64-binary" ]