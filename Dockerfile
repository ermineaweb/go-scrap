FROM golang:1.18 as builder

ARG SERVICE_NAME
ENV APP_NAME=$SERVICE_NAME

WORKDIR /build
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=0 go build -o ./bin/$APP_NAME ./cmd/$APP_NAME

FROM alpine:3.15

RUN apk add --update --no-cache 	bash 	nano 	jq 	tcpdump 	curl 	kafkacat

ARG SERVICE_NAME
ENV APP_NAME=$SERVICE_NAME

WORKDIR /app
COPY --from=builder /build/bin/$APP_NAME ./$APP_NAME
COPY ./entrypoint.sh ./entrypoint.sh 
RUN chmod +x entrypoint.sh

COPY ./scripts/ ./
RUN chmod +x *.sh

ENTRYPOINT ["/app/entrypoint.sh"]
