FROM golang:alpine


WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary

EXPOSE 3000

ENTRYPOINT ["/app/binary"]