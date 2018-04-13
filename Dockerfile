FROM golang:1.9.0-alpine3.6 AS build-env
RUN apk add --no-cache git
ADD . /src
RUN cd /src && go-wrapper download \
    && CGO_ENABLED=0 GOOS=linux \
    go build -a -ldflags '-extldflags "-static"' -o hercules;

# final stage
FROM alpine:3.6
WORKDIR /
RUN apk --no-cache update && \
    apk --no-cache add ca-certificates
COPY --from=build-env /src/hercules /hercules
CMD ["/hercules"]
