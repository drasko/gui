FROM golang:1.15-alpine AS builder

WORKDIR /go/src/github.com/drasko/gui
COPY . .
RUN make $SVC_NAME \
    && mv build/gui /exe

FROM scratch
COPY --from=builder /exe /
ENTRYPOINT ["/exe"]
