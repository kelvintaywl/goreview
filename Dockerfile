FROM golang:1.9-alpine AS goreview-build

RUN apk --update add make git

WORKDIR /go/src/github.com/kelvintaywl/goreview

ADD . /go/src/github.com/kelvintaywl/goreview
RUN make init build


FROM alpine:latest

COPY --from=goreview-build /go/src/github.com/kelvintaywl/goreview/goreview /goreview

EXPOSE 9999

ENTRYPOINT ["/goreview"]
