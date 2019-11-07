FROM golang:alpine as build-env
WORKDIR /usr/bin/desafio-graph
ADD . /usr/bin/desafio-graph
RUN go build -ldflags "-s -w" -o server

# ------ 

FROM alpine
WORKDIR /usr/bin
COPY --from=build-env /usr/bin/desafio-graph/server ./
ENTRYPOINT /usr/bin/server

EXPOSE 8080