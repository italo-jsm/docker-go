FROM golang:alpine as build-env
WORKDIR /
COPY . . 
RUN  go build -o /usr/bin/palestra
FROM alpine
COPY --from=build-env /usr/bin/palestra /usr/bin
ENTRYPOINT [ "/usr/bin/palestra" ]