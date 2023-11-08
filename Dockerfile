FROM cgr.dev/chainguard/go AS builder

COPY . /app
RUN cd /app && go build -o bucket-website-app

FROM cgr.dev/chainguard/glibc-dynamic
COPY --from=builder /app/bucket-website-app /usr/bin
CMD ["/usr/bin/bucket-website-app"]
