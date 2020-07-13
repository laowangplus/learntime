FROM golang

COPY . /

WORKDIR /

RUN go run main.go

EXPOSE 8888