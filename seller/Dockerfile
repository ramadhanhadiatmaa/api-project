FROM golang:1.22.8 AS build

WORKDIR /app

COPY go.mod go.sum ./
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o seller

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/seller .

EXPOSE 8711

CMD [ "./master" ]