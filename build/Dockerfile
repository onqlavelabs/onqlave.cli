FROM golang:1.18 as build
LABEL org.opencontainers.image.source="https://github.com/onqlavelabs/onqlave.all"
LABEL org.opencontainers.image.description="This is onqlave cli container, used to interact with Onqlave platform. For more information, read our documentation at https://docs.onqlave.com"

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go mod tidy

RUN CGO_ENABLED=0 go build -o /go/src/app/build/cli/linux/onqlave /go/src/app/cmd/cli

FROM bash:5.2

ARG CLI_ENV="prod"

COPY --from=build /go/src/app/build/cli/linux/onqlave /usr/local/bin/onqlave
ENV ENV $CLI_ENV
ENV TERM xterm-256color

CMD ["bash", "-c", "onqlave config init && exec bash"]
