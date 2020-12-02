FROM golang:1.14

RUN go get -v github.com/gin-gonic/gin
RUN go get -v github.com/gin-contrib/cors
RUN go get -v github.com/globalsign/mgo/bson
RUN go get -v github.com/go-bongo/bongo

RUN mkdir /go/src/openbankingcrawler/ 
ADD . /go/src/openbankingcrawler
WORKDIR /go/src/openbankingcrawler

RUN go install

RUN go build main.go 

EXPOSE 8090

CMD ["./main"]