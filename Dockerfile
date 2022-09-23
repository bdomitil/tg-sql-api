FROM golang:1.18

WORKDIR /usr/src/app


##  REQUIRS
## TG_DB_PASSWD - database password for db
## TG_DB_USER - database user 
## TG_DB - database name
## TG_DB_HOST - database host address

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/tg-cache .

CMD ["tg-cache"]