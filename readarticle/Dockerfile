ARG GO_VERSION=1.12

FROM golang:$GO_VERSION as builder

ARG ssh_prv_key
ARG ssh_pub_key

RUN mkdir -p /root/.ssh && \
    chmod 0700 /root/.ssh && \
    ssh-keyscan github.com > /root/.ssh/known_hosts

# Add the keys and set permissions
RUN echo "$ssh_prv_key" > /root/.ssh/id_rsa && \
    echo "$ssh_pub_key" > /root/.ssh/id_rsa.pub && \
    chmod 600 /root/.ssh/id_rsa && \
    chmod 600 /root/.ssh/id_rsa.pub

# Use ssh instead of https
RUN git config --global url."git@github.com:".insteadOf "https://github.com/"

# modules
ENV GO111MODULE=on
WORKDIR $GOPATH/src/readarticle
ADD . $GOPATH/src/rearticle

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
COPY readarticle.sh /bin
RUN chmod 777 /bin/readarticle.sh

# build time
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /go/bin/readarticle cmd/main.go

# RUN apt-get -y install  sysfsutils apt-utils
RUN echo exit 0 > /usr/sbin/policy-rc.d
# install and run redis (in the background)
RUN apt-get update && apt-get -y install redis-server
# RUN service redis-server start --bind 127.0.0.1
RUN ls /go/bin

# run options
ENV PORT_NUM=5100
EXPOSE 5100
ENTRYPOINT ["readarticle.sh"]
