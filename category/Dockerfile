FROM golang:1.13

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /go/src
ADD app .
ADD script .
RUN chmod a+x entrypoint.sh

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o service ./service

EXPOSE 8000
ENTRYPOINT ["./entrypoint.sh"]
