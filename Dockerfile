FROM golang:latest

WORKDIR "/app"

COPY ["./go.mod", "./go.sum", "./"]
ADD . .

EXPOSE 9999

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin ./

CMD [ "./bin" ]