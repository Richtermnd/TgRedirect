FROM golang:latest

WORKDIR "/app"

COPY ["./go.mod", "./"]
ADD . .

EXPOSE 9999

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin ./

CMD [ "./bin" ]