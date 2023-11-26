FROM golang

COPY main.go /Bob/main.go
WORKDIR /Bob/
RUN go build -o main main.go
ENV ALICE_URL="http://alice:8080"
ENV EVE_URL="http://eve:8080"
CMD ["./main"]
EXPOSE 8080
