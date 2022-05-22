FROM golang:alpine AS builder

ARG auth_token
ARG db_connection_string

ENV auth_token=$auth_token
ENV db_connection_string=$db_connection_string

# Install git
RUN apk update && apk add --no-cache git && apk add openssh && apk add sed && apk add curl

# Create user
ENV USER=builduser
ENV UID=10001 

# Create an unprivileged user
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

WORKDIR $GOPATH/src/service
COPY . .
COPY ./etc/config.yaml /etc/config.yaml

# create ssh directory
RUN mkdir ~/.ssh
RUN touch ~/.ssh/known_hosts
RUN ssh-keyscan -t rsa github.com >> ~/.ssh/known_hosts

RUN sed -r "s/^(\s*)(image\s*:\s*dbtemplate\s*$)/\1image: ${db_connection_string}/" /etc/config.yaml

RUN curl --create-dirs -o /etc/root.crt -O https://cockroachlabs.cloud/clusters/e437f606-804a-4bc6-bbc7-ba1470eb1dd4/cert

# allow private repo pull
RUN git config --global url."https://${auth_token}:x-oauth-basic@github.com/".insteadOf "https://github.com/"

# Fetch dependencies
RUN go get -d -v

# Build
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/service

FROM scratch

# Grab user files from previous stage
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/config.yaml /etc/config.yaml
COPY --from=builder /go/bin/service /go/bin/service
COPY --from=builder /etc/root.crt /.postgresql/root.crt 

# Exec as an unprivileged user
USER builduser:builduser

EXPOSE 8000

ENTRYPOINT [ "/go/bin/service" ]