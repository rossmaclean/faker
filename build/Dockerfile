FROM golang:alpine as go-build
WORKDIR /app
ADD ./cmd/faker ./cmd/faker
ADD ./configs ./configs
ADD ./internal ./internal
ADD ./go.mod .
ADD ./go.sum .
RUN #go test -v ./...
RUN go build -o faker /app/cmd/faker

FROM golang:alpine
WORKDIR /app/code
COPY --from=go-build /app/faker ./api/
ADD ./configs/*.properties ./configs/

RUN adduser -S -D -H -h /app appuser
USER appuser
EXPOSE 8000
CMD ["./api/faker"]