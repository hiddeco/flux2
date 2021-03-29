FROM alpine:3.13 as builder

RUN apk add --no-cache ca-certificates curl

ARG ARCH
ARG KUBECTL_VER

RUN kubectl_ver=${KUBECTL_VER:-1.20.4} && \
    arch=${ARCH:-linux/amd64} && \
    curl -sL https://storage.googleapis.com/kubernetes-release/release/v${kubectl_ver}/bin/${arch}/kubectl \
    -o /usr/local/bin/kubectl && chmod +x /usr/local/bin/kubectl && \
    kubectl version --client=true

FROM alpine:3.13 as flux-cli

LABEL org.opencontainers.image.source="https://github.com/fluxcd/flux2"

# Create minimal nsswitch.conf file to prioritize the usage of /etc/hosts over DNS queries.
# https://github.com/gliderlabs/docker-alpine/issues/367#issuecomment-354316460
RUN [ ! -e /etc/nsswitch.conf ] && echo 'hosts: files dns' > /etc/nsswitch.conf

RUN apk add --no-cache ca-certificates

COPY --from=builder /usr/local/bin/kubectl /usr/local/bin/
COPY --chmod=755 flux /usr/local/bin/

ENTRYPOINT [ "flux" ]
