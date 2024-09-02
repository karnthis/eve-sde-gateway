FROM golang:1.22 AS build
WORKDIR /src
COPY ./ ./
ENV CGO_ENABLED=0 GOOS=linux GOARCH=arm64
RUN go build -o /src/sde
RUN pwd && ls -hal

FROM scratch
WORKDIR /bin
COPY --from=build /src/sde /bin/sde
CMD ["/bin/sde"]