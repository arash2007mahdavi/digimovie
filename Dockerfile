FROM golang:1.23.6-alpine
COPY /src /app
WORKDIR /app
RUN go build -o app main.go
EXPOSE 1386
CMD [ "./app" ]
