FROM golang:1.21

LABEL maintainer="you@example.com"

RUN apt-get update && \
    apt-get install -y curl unzip && \
    apt-get clean

ENV GOFUMPT_VERSION=0.4.0 \
    GCI_VERSION=0.10.1 \
    GOLANGCI_LINT_VERSION=1.57.2 \
    GOBIN=/usr/local/bin

RUN curl -L "https://github.com/mvdan/gofumpt/releases/download/v${GOFUMPT_VERSION}/gofumpt_v${GOFUMPT_VERSION}_linux_amd64" -o /usr/local/bin/gofumpt && \
    go install github.com/daixiang0/gci@v${GCI_VERSION} && \
    curl -L "https://github.com/golangci/golangci-lint/releases/download/v${GOLANGCI_LINT_VERSION}/golangci-lint-${GOLANGCI_LINT_VERSION}-linux-amd64.tar.gz" -o /tmp/lint.tar.gz && \
    cd /tmp && \
    tar -xzf lint.tar.gz && mv golangci-lint-*/golangci-lint /usr/local/bin/ && \
    chmod +x /usr/local/bin/* && \
    rm -rf /tmp/*