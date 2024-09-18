FROM golang:1.23 AS build
ADD . /src
WORKDIR /src
RUN go build -v -o app

FROM alpine:3.20.3
EXPOSE 8080
COPY --from=build /src/app /usr/local/bin/app
RUN chmod +x /usr/local/bin/app
CMD ["app"]