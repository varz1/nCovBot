FROM golang as gobuilder

WORKDIR /build

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /nCovBot


FROM scratch

COPY --from=gobuilder /list.json /
COPY --from=gobuilder /nCovBot /

CMD ["/nCovBot"]