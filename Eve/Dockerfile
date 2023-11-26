FROM golang

COPY main.go /Eve/main.go
WORKDIR /Eve/
RUN go build -o main main.go
CMD ["./main"]
EXPOSE 8080
