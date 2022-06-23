FROM golang:1.18-alpine

ADD . /go/src/myapp
WORKDIR /go/src/myapp

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go install

COPY *.go ./

RUN go build -o /docker-go-slackhook

CMD ["/docker-go-slackhook"]