FROM golang:1.18-alpine as build
WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./api .

FROM alpine:3.9
RUN apk add ca-certificates
WORKDIR /app
COPY --from=build /app/api .
COPY --from=build /app/.env .      


EXPOSE 5000

CMD ["./api"]