FROM golang:alpine

WORKDIR /airbnb-clone

COPY . /airbnb-clone

COPY go.mod go.sum ./

RUN go mod download

RUN go build -o main main.go

EXPOSE 8080

# RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-airbnb

CMD ["main"]