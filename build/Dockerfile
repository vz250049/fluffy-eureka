FROM golang:1.16-alpine AS build

WORKDIR /cmd

COPY . ./

RUN go mod download

#RUN CGO_ENABLED=0 go test -v /tmp/app/utils/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sample ./cmd/main.go

FROM alpine:latest

USER 1001:1001

WORKDIR /cmd

COPY --from=build /cmd/sample /cmd

ENTRYPOINT [ "/cmd/sample"]