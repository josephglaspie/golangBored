FROM golang:1.17

ENV APP_NAME myproject
ENV PORT 8080

COPY . /go/src/${APP_NAME}
WORKDIR /go/src/${APP_NAME}

RUN go get ./
RUN go build -o ${APP_NAME}
EXPOSE ${PORT}

CMD ./${APP_NAME}

