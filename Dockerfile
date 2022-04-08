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

ARG SERVICE_NAME
ENV APP_NAME=$SERVICE_NAME

WORKDIR /app
COPY --from=builder /build/bin/$APP_NAME ./$APP_NAME
COPY ./entrypoint.sh ./entrypoint.sh 
RUN chmod +x entrypoint.sh

ENTRYPOINT ["/app/entrypoint.sh"]
