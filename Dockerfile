FROM golang:1.16 as build

COPY . /src

WORKDIR /src

RUN go mod download

RUN go build -o main

FROM ubuntu

COPY --from=build src/main main
COPY --from=build src/config/config.yaml config.yaml

EXPOSE 8080

CMD ./main
