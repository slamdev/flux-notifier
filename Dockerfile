FROM alpine

COPY flux-notifier /usr/bin/

ENTRYPOINT ["flux-notifier", "listen"]
