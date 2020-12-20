FROM golang:1.15.5

WORKDIR /usr/app
COPY ../.. .

RUN make build

CMD ["./target/bot"]
