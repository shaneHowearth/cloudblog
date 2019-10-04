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
WORKDIR $GOPATH/src/REST
ADD . $GOPATH/src/REST

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# build time
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /go/bin/REST cmd/main.go cmd/routes.go

# run options
ENV PORT_NUM=5055
EXPOSE 5055
ENTRYPOINT ["REST"]
