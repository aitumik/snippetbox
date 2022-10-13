FROM golang:1.17.5-alpine AS development

MAINTAINER aitumik@protonmail.com

# Set the neccessary environment variables
ENV G0111MODULE=on
ENV CGO_ENABLED 0
ENV GOPATH /go
ENV GOCACHE /go-build

# Create context directory
RUN mkdir /app

# Change the current directory to present directory
WORKDIR /app

# Copy everything from the current dir to present dir
COPY . /app

RUN --mount=type=cache,target=/go/pkg/mod/cache \
    go mod download

RUN --mount=type=cache,target=/go/pkg/mod/cache \
      --mount=type=cache,target=/go-build \
      go build -o sniper cmd/web/*

# Install gcc dependencies
#RUN apk add git alpine-sdk build-base gcc

CMD ["./sniper"]

FROM development AS production

COPY --from=gloursdocker/docker / /

CMD ["go", "run", "cmd/web/*"]

FROM scratch
COPY --from=development /app/sniper /usr/bin/sniper
CMD ["/usr/bin/sniper"]

