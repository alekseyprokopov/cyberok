FROM golang

RUN go version
ENV GOPATH=/

COPY ./ ./

#psql
RUN apt-get update
RUN apt-get -y install postgresql-client
RUN chmod +x wait-for-postgres.sh
RUN


#build
RUN go mod download
RUN go build -o cyberok ./cmd/cyberok/main.go

CMD ["./cyberok"]