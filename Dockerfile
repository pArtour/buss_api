FROM golang:1.18.2-alpine as build
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
RUN ls -al

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/main .
COPY --from=build /app/migrate ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration
COPY db_tables .

EXPOSE 8000
CMD ["/app/main"]
ENTRYPOINT ["/app/start.sh"]