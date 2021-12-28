FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go build -o random-quotes-backend

EXPOSE 3000

CMD ./random-quotes-backend