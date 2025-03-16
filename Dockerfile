FROM golang:1.23.6-alpine
COPY . /app
WORKDIR /app
RUN go build -o app main.go
EXPOSE 1386
CMD [ "./app" ]

# docker run -d --network db-network --name pgdb1 -p 5432:5432 --env POSTGRES_USER=arashmahdavi --env POSTGRES_PASSWORD=@rash2007 postgres
# docker run -d --name pgadmin --network db-network -p 8090:80 --env PGADMIN_DEFAULT_EMAIL=arashmahdavi2007@gmail.com --env PGADMIN_DEFAULT_PASSWORD=@rash2007 dpage/pgadmin4