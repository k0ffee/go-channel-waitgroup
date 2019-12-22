FROM golang:1.13 as builder
WORKDIR /var/tmp
USER daemon
COPY Makefile .
COPY backend.go .
COPY backend_test.go .
ENV GOCACHE=/var/tmp/.gocache
RUN make
RUN make test

FROM scratch
WORKDIR /
COPY --from=builder /var/tmp/backend .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER 1
ENTRYPOINT ["/backend"]
