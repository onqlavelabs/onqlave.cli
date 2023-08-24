FROM golang:1.18 as build
LABEL org.opencontainers.image.source="https://github.com/onqlavelabs/onqlave.cli"
LABEL org.opencontainers.image.description="This is onqlave cli container, used to interact with Onqlave platform. For more information, read our documentation at https://docs.onqlave.com"

ARG CLI_VERSION="alpha"

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go mod tidy

RUN CGO_ENABLED=0 go build -ldflags="-X 'github.com/onqlavelabs/onqlave.cli/cmd.Version=$CLI_VERSION'" -o /go/src/app/bin/onqlave /go/src/app

FROM bash:5.2

RUN mkdir -p /root/.config/onqlave
COPY --from=build /go/src/app/bin/onqlave /usr/local/bin/onqlave
ENV TERM xterm-256color

CMD ["bash", "-c", "onqlave && exec bash"]
