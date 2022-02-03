FROM golang:alpine3.15 as compiler
WORKDIR /go-tour
RUN apk update && apk add upx
COPY ./ /go-tour/
RUN ./build.sh

FROM alpine:3.15
WORKDIR /go-tour
COPY --from=compiler ./ ./
CMD ["./run.sh"]
