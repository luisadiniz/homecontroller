FROM golang:alpine as build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . ./
RUN go build -o /homecontroller

FROM alpine
WORKDIR /
COPY --from=build /homecontroller /homecontroller
EXPOSE 80
ENTRYPOINT [ "/homecontroller" ]