FROM golang:1.18.2
ENV GO111MODULE=on
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download


#COPY *.go ./
COPY . .
RUN go build -o /docker-stresstest

EXPOSE 8080

CMD [ "/docker-stresstest" ]