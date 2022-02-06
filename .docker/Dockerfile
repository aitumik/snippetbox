FROM golang:1.17.5-alpine
MAINTAINER aitumik@protonmail.com

# Set the neccessary environment variables
ENV G0111MODULE=on

# Change the current directory to present directory
WORKDIR /build

# Copy everything from the current dir to present dir
COPY . .

# Install gcc dependencies
RUN apk --no-cache add git alpine-sdk build-base gcc

# Expose port 4000
EXPOSE 4000

RUN go build -o sniper cmd/web/*

CMD ["./sniper"]


