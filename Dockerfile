
FROM golang:alpine AS gobuilder
RUN apk add git
WORKDIR /go/src/dataCollectionTree
ADD . /go/src/dataCollectionTree
RUN cd /go/src/dataCollectionTree && go build -o main
FROM alpine:latest 
WORKDIR /root/
COPY --from=gobuilder /go/src/dataCollectionTree/main .
COPY --from=gobuilder /go/src/dataCollectionTree/vars.env .
RUN source vars.env
CMD ["./main"]