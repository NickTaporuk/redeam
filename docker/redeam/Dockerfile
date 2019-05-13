# I use 2 stage dockerfile build for save space in disk

# default golang version
# you can change golang version if it will need
ARG GOLANG_VERSION=1.12.4-stretch

# default alpine container version
ARG ALPINE_CONTAINER_VERSION=3.7

FROM golang:${GOLANG_VERSION} as scratch-base

ENV PACKAGE_NAME github.com/NickTaporuk/redeam
# application name
ENV COMPAILED_APP_NAME redeam-test-task

RUN go get -d -v -u ${PACKAGE_NAME}

WORKDIR ${GOPATH}/src/${PACKAGE_NAME}

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o /$COMPAILED_APP_NAME main.go

# end of first stage, beginning of the second one
FROM alpine:${ALPINE_CONTAINER_VERSION}

ENV COMPAILED_APP_NAME redeam-test-task

COPY --from=scratch-base /$COMPAILED_APP_NAME /$COMPAILED_APP_NAME

RUN chmod +x ${COMPAILED_APP_NAME}

CMD /$COMPAILED_APP_NAME
