FROM golang:1.24

WORKDIR /usr/bin/app


COPY go.mod .
CMD ["go", "mod", "tidy"]

COPY . .
EXPOSE 8000