ARG build_image=golang:1.18.3-alpine3.16

FROM ${build_image} as build_stage

ENV GO111MODULE=on

# Application env
ENV APP_VERSION=0.1.1 \
    TIMEZONE=Asia/Almaty \

    # Database
    DB_DRIVER=postgres \
    DB_HOST=localhost \
    DB_PORT=5432 \
    DB_USER=paimon \
    DB_PASS=paimon \
    DB_NAME=paimoncookies \
    DB_SSL_MODE=disable \

    # Bot
    BOT_TIMEOUT=60 \
    BOT_DEBUG=true \
    BOT_TEMPLATE_PATH=internal/commander/template/ \

    # API
    API_DEBUG=true \
    API_APPKEY=test \
    API_PORT=8087

WORKDIR /code

COPY go.mod go.sum ./

RUN go mod download -x && go mod verify

COPY . .

RUN go mod vendor -v;
RUN go build -v -mod=vendor -o /usr/local/bin/paimoncookies ./cmd/paimoncookies

EXPOSE $API_PORT

CMD ["/usr/local/bin/paimoncookies"]

LABEL version=$APP_VERSION \
      maintainer="Anna Vassilenko <esabril.ch@gmail.com>" \
      name="Paimon Cookies Bot" \
      description="Useful information about Genshin Impact game"