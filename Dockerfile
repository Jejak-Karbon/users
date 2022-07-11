FROM golang:alpine

WORKDIR /var/www/html/apps/users
COPY . /var/www/html/apps/users

RUN go build -o main .

CMD ["/var/www/html/apps/users/main"]