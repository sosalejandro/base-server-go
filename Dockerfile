FROM golang:latest AS build-env

# Allow Go to retrieve the dpendencies for the build step
RUN apk add --no-cache git

# Secure against running as root
RUN adduser -D -u 10000 alejandro
RUN mkdir /web-server/ && chown alejandro /web-server/
USER alejandro

WORKDIR /web-server/
ADD . /web-server/

# Compile the binary, we don't want to run the cgo resolver
RUN CG0_ENABLED=0 go build -o /web-server/app .

# final stage
FROM alphine:latest

# Secure against running as root
RUN adduser -D -u 10000 alejandro
USER alejandro

WORKDIR /
COPY --from=build-env /web-server/certs/docker.localhost.* /
COPY --from=build-env /web-server/app /

EXPOSE 4444

CMD ["/app"]