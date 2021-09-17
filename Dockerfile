FROM golang
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main
CMD [ "./main" ]