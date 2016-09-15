
FROM golang:1.7rc6
MAINTAINER Ernesto Alejo <ernesto@altiplaconsulting.com>

RUN apt-get update && \
    apt-get install -y pkg-config cmake build-essential

COPY tools/install-libgit2.sh /opt/
RUN /opt/install-libgit2.sh

WORKDIR /go/src/github.com/altipla-consulting/rls
CMD go install ./cmd/rls && mv /go/bin/rls /opt/build
