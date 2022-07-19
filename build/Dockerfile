FROM golang:alpine as go-build
WORKDIR /app
ADD ./cmd/faker ./cmd/faker
ADD ./internal ./internal
ADD ./go.mod .
ADD ./go.sum .
RUN #go test -v ./...
RUN go build -o faker /app/cmd/faker

FROM cloudron/base:3.2.0@sha256:ba1d566164a67c266782545ea9809dc611c4152e27686fd14060332dd88263ea

RUN mkdir -p /app/code/api
WORKDIR /app/code

COPY --from=go-build /app/faker /app/code/api/faker
ADD ./configs/*.properties /app/code/configs/
COPY start.sh /app/code/

RUN chown -R cloudron:cloudron /app/code

CMD [ "/app/code/start.sh" ]