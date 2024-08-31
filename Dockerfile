FROM golang:1.22.5-alpine3.20 AS build

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go test --cover -v ./...
RUN go build -v -o groupie

FROM alpine:latest

LABEL authors="S.Cointin, A.Nassuif, M.Soumare"\
      description="Groupie Tracker is a simple website that tracks artists info based on a JSON API"\
      licence="GNU GPL V3.0-or-later"\
      maintainer="See author label"\
      contact="Github profile"

WORKDIR /app
COPY --from=build /app/groupie /app/groupie
COPY --from=build /app/src /app/src

EXPOSE 5826

CMD ["/app/groupie"]