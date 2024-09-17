FROM golang:1.23 AS build
ADD . /src
WORKDIR /src
RUN go build -v -o learning-go

FROM alpine:3.20.3
COPY --from=build /src/learning-go /usr/local/bin/learning-go
RUN chmod +x /usr/local/bin/learning-go
CMD ["./usr/local/bin/learning-go"]