FROM golang:1.17.6-bullseye

ARG DEBIAN_VERSION=11

RUN wget -qO- https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > microsoft.asc.gpg \
    && mv microsoft.asc.gpg /etc/apt/trusted.gpg.d/ \
    && wget -q https://packages.microsoft.com/config/debian/${DEBIAN_VERSION}/prod.list \
    && mv prod.list /etc/apt/sources.list.d/microsoft-prod.list \
    && chown root:root /etc/apt/trusted.gpg.d/microsoft.asc.gpg \
    && chown root:root /etc/apt/sources.list.d/microsoft-prod.list

RUN apt-get update && apt-get -y install --no-install-recommends azure-functions-core-tools-3

RUN GO111MODULE=on go get -v \
    golang.org/x/tools/gopls@latest \
    golang.org/x/lint/golint@latest \
    github.com/go-delve/delve/cmd/dlv@latest \
    2>&1

RUN GO111MODULE=on go get -v \
    github.com/jstemmer/go-junit-report@latest \
    github.com/axw/gocov/gocov@latest \
    github.com/AlekSi/gocov-xml \
    github.com/wadey/gocovmerge \
    2>&1

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.43.0

RUN echo 'export DOTNET_SYSTEM_GLOBALIZATION_INVARIANT=1' >> /root/.bashrc
