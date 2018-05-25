FROM golang:1.10.2

EXPOSE 1323

WORKDIR /go/src/app

COPY . .

WORKDIR /go/src/app/internal/app
RUN go install

CMD ["app"]

