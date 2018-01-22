FROM golang:1.9-alpine AS goreview-build

ARG GITHUB_ACCESS_TOKEN

RUN echo "machine github.com login $GITHUB_ACCESS_TOKEN" > ~/.netrc
RUN apk --update add make git

WORKDIR /go/src/github.com/kelvintaywl/goreview

ADD . /go/src/github.com/kelvintaywl/goreview
RUN make init build


FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY ./cert.pem /usr/local/share/ca-certificates/mycert.crt
RUN update-ca-certificates
COPY --from=goreview-build /go/src/github.com/kelvintaywl/goreview/goreview /goreview

EXPOSE 9999

ENTRYPOINT ["/goreview"]
