FROM golang:1.19-alpine AS build

WORKDIR /go/src/rest

COPY . .

RUN CGO_ENABLED=0 go build -o /go/bin/rest

FROM alpine:3.17

COPY --from=build /go/bin/rest /go/bin/rest

ENTRYPOINT ["/go/bin/rest"]
# trigger pipeline