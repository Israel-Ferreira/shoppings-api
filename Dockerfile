FROM golang:1.23.2-alpine AS build

WORKDIR /usr/app

COPY . /usr/app

RUN CGO_ENABLED=0 GOOS=linux go build -o shoppings-api main.go


FROM scratch

COPY --from=build /usr/app/shoppings-api  .

EXPOSE 4000

CMD ["./shoppings-api"]