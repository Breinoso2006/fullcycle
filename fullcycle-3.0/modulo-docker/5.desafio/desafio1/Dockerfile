FROM golang:alpine as builder

WORKDIR /go/src

COPY /go .

RUN CGO_ENABLED=0 GOOS=linux go build -o exec fc.go

FROM scratch

COPY --from=builder /go/src .

ENTRYPOINT [ "./exec" ]