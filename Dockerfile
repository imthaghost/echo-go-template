FROM ubuntu:bionic as builder

ENV GOLANG_VERSION 1.17.1
ENV GOLANGCI_LINT_VERSION 1.39.0
ENV GOLANG_SRC_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOPATH=/build/go
ENV PATH=/usr/local/go/bin:/build/go/bin:$PATH

# Helps identify build image id when fetching test coverage report.
LABEL build=imthaghost/echo-template

# Update
RUN  apt-get update \
  && apt-get install -y wget \
  && apt-get install make \
  && rm -rf /var/lib/apt/lists/* \

RUN sudo apt-get install build-essential

# we download go here and extract it to usr/local
RUN wget "$GOLANG_SRC_URL" && \
      tar -C /usr/local -xzf go$GOLANG_VERSION.linux-amd64.tar.gz && \
      rm go$GOLANG_VERSION.linux-amd64.tar.gz


# download and install golangci lint software
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v$GOLANGCI_LINT_VERSION

WORKDIR /build/go/src/github.com/imthaghost/echo-template
COPY . .

# validate test and lint is good
# RUN make test && make lint

# build the binary
RUN make echo-template

# Copy binary into seperate contianer. This creates a clean space to run the application from.
FROM ubuntu:bionic

WORKDIR /root/
COPY --from=builder /build/go/src/github.com/imthaghost/echo-template/echo-template .

EXPOSE 8080

RUN apt-get update; apt-get install -y ca-certificates

CMD ["./echo-template"]
