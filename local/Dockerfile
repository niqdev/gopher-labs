FROM node:alpine3.17

RUN \
  apk upgrade --update && \
  apk add --no-cache \
    bash \
    curl \
    jq \
    aws-cli

CMD ["/bin/bash"]
