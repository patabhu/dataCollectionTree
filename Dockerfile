
FROM golang:alpine AS gobuilder
RUN apk add git
ADD . /go/src/dataCollectionTree
WORKDIR /go/src/dataCollectionTree
RUN go build -o main
FROM alpine:latest 
WORKDIR /root/
COPY --from=gobuilder /go/src/dataCollectionTree/main .
RUN source /go/src/dataCollectionTree/.env
CMD ["./main"]