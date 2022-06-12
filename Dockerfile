ARG build_image=golang:1.18.3-alpine3.16

FROM ${build_image} as build_stage

ENV GO111MODULE=on

WORKDIR /code

COPY go.mod go.sum ./

RUN go mod download -x && go mod verify

COPY . .

RUN go mod vendor -v;
RUN go build -v -mod=vendor -o /usr/local/bin/paimoncookies ./cmd/paimoncookies

EXPOSE 8087

CMD ["/usr/local/bin/paimoncookies"]

LABEL version=0.0.1 \
      maintainer="Anna Vassilenko <esabril.ch@gmail.com>" \
      name="Paimon Cookies Bot" \
      description="Useful information about Genshin Impact game"