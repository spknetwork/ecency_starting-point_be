# syntax=docker/dockerfile:1

FROM golang:1.19

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
COPY . ./

RUN go build -o run ./
RUN chmod +x ./run
ENV PORT 8000
EXPOSE $PORT
CMD ["./run"]