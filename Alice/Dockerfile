FROM golang

COPY main.go /Alice/main.go
WORKDIR /Alice/
RUN go build -o main main.go
CMD ["./main"]
EXPOSE 8080
