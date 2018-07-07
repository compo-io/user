FROM golang:1.10
WORKDIR /go/src/github.com/compo-io/user
ADD ./ .
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -o user.bin -ldflags '-s -w -extldflags "-static"' main.go

FROM scratch
COPY --from=0 /go/src/github.com/compo-io/user/user.bin /
CMD ["/user.bin"]