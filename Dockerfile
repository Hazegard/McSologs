FROM alpine:latest as build

RUN apk add go

COPY . /usr/local/src/

WORKDIR /usr/local/src/

RUN go build -o /usr/local/src/mcsologs .




FROM alpine:latest

COPY --from=build /usr/local/src/mcsologs /usr/local/bin/mcsologs

ENTRYPOINT [ "/usr/local/bin/mcsologs" ]