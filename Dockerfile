ARG build_image=golang:1.18.3-alpine3.16

FROM ${build_image} as build_stage

ENV GO111MODULE=on
ARG tag_version

# Application env
ENV PCOOKIES_APP_VERSION=$tag_version \
    PCOOKIES_PCOOKIES_TIMEZONE=Asia/Almaty \

    # Database
    PCOOKIES_DB_DRIVER=postgres \
    PCOOKIES_DB_HOST=localhost \
    PCOOKIES_DB_PORT=5432 \
    PCOOKIES_DB_USER=paimon \
    PCOOKIES_DB_PASS=paimon \
    PCOOKIES_DB_NAME=paimoncookies \
    PCOOKIES_DB_SSL_MODE=disable \

    # Bot
    PCOOKIES_BOT_TIMEOUT=60 \
    PCOOKIES_BOT_DEBUG=true \
    PCOOKIES_BOT_TEMPLATE_PATH=internal/commander/template/ \

    # API
    PCOOKIES_API_DEBUG=true \
    PCOOKIES_API_APPKEY=test \
    PCOOKIES_API_PORT=8087

WORKDIR /code

COPY go.mod go.sum ./

RUN go mod download -x && go mod verify

COPY . .

RUN go mod vendor -v;
RUN go build -v -mod=vendor -o /usr/local/bin/paimoncookies ./cmd/paimoncookies

EXPOSE $PCOOKIES_API_PORT

CMD ["/usr/local/bin/paimoncookies"]

LABEL version=$PCOOKIES_APP_VERSION \
      maintainer="Anna Vassilenko <esabril.ch@gmail.com>" \
      name="Paimon Cookies Bot" \
      description="Useful information about Genshin Impact game"