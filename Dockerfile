FROM golang:1.17.5-alpine AS development

MAINTAINER aitumik@protonmail.com

# Set the neccessary environment variables
ENV G0111MODULE=on

# Create context directory
RUN mkdir /build

# Change the current directory to present directory
WORKDIR /build

# Copy everything from the current dir to present dir
COPY . /build
COPY tls ./build

# Install gcc dependencies
RUN apk add git alpine-sdk build-base gcc


FROM development AS production

# Expose port 4000
EXPOSE 4000

RUN go build -o sniper cmd/web/*

CMD ["./sniper"]


