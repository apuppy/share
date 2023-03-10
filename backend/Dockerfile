FROM golang:1.18
WORKDIR /go/src/github.com/apuppy/share/backend
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct && CGO_ENABLED=0 go build -v -a -installsuffix cgo -o share-backend .

FROM debian:bullseye
# RUN apk --no-cache add ca-certificates
WORKDIR /www
COPY --from=0 /go/src/github.com/apuppy/share/backend/share-backend ./
CMD ["./share-backend"]
