FROM golang:1.22.3-alpine

WORKDIR /application

COPY . /application

RUN go build -o output

CMD [ "./output" ]
