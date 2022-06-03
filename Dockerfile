FROM golang:alpine

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

RUN mkdir /app
WORKDIR /app

COPY . .
COPY .env .

RUN go get -d -v ./... && go install -v ./...

RUN cd src/ && go build -o /build

EXPOSE ${API_PORT}

CMD [ "/build" ]