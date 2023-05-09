FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o dist

EXPOSE 5000

ENTRYPOINT [ "./dist" ]