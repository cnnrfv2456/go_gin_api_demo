FROM golang:1.19.1
ENV GO111MODULE=on
WORKDIR /api
COPY  . .
# RUN go mod download
# RUN go mod vendor
RUN go mod tidy
RUN go build -o main .
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2
RUN go install github.com/swaggo/swag/cmd/swag@v1.8.6
# CMD ["./main"]
