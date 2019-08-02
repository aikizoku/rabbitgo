FROM golang:1.12-alpine AS build_base

RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR /go/src/github.com/aikizoku/rabbitgo

ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .

RUN go mod download

FROM build_base AS server_builder
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go install -a -tags netgo -ldflags '-w -extldflags "-static"' .

FROM alpine AS server

RUN apk add ca-certificates
COPY --from=server_builder /go/bin/rabbitgo /bin/rabbitgo
COPY --from=server_builder /go/src/github.com/aikizoku/rabbitgo/.env /go/src/github.com/aikizoku/rabbitgo/.env
COPY --from=server_builder /go/src/github.com/aikizoku/rabbitgo/serviceAccount.json /go/src/github.com/aikizoku/rabbitgo/serviceAccount.json

WORKDIR /go/src/github.com/aikizoku/rabbitgo

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT ["/bin/rabbitgo"]