FROM golang:alpine
EXPOSE 8000:8000
RUN mkdir /app
WORKDIR /app

COPY . /app
COPY .env /app

RUN go get -d ./...

RUN go build -o app ./cmd/api/main.go 

CMD [ "./app" ]