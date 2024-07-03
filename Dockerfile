FROM golang:1.22 AS build
WORKDIR /app
COPY go.mod go.sum main.go ./
RUN --mount=type=cache,target=/go/pkg/mod/ go mod download -x
COPY ./app ./app
RUN --mount=type=cache,target=/go/pkg/mod/ CGO_ENABLED=0 go build -tags lambda.norpc -o main .

FROM public.ecr.aws/lambda/provided:al2023
COPY --from=build /app/main ./main
ENTRYPOINT [ "./main" ]
