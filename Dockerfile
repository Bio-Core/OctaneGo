FROM golang:1.8

WORKDIR /go/src/github.com/Bio-core/OctaneGo
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN make build

CMD ./bin/octanego