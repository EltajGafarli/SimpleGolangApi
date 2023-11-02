FROM golang

WORKDIR /app

COPY . /app/

RUN go mod download

EXPOSE 3000

RUN go build


CMD ["./gobank"]
