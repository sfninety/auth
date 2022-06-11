FROM golang:alpine AS builder

ARG auth_token
ARG db_connection_string
ARG db_cert
ARG jwt_signing_key

ENV auth_token=$auth_token
ENV db_connection_string=$db_connection_string
ENV db_cert=$db_cert
ENV jwt_signing_key=$jwt_signing_key

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


RUN curl --create-dirs -o /etc/root.crt -O https://cockroachlabs.cloud/clusters/${db_cert}/cert

# allow private repo pull
RUN git config --global url."https://${auth_token}:x-oauth-basic@github.com/".insteadOf "https://github.com/"

# Fetch dependencies
RUN go get -d -v

# Replace the templated database connection string
RUN sed -ri "s#dbtemplate#'${db_connection_string}'#" /etc/config.yaml
RUN sed -ri "s#jwtsigningkeytemplate#'${jwt_signing_key}'#" /etc/config.yaml

# Build
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o /go/bin/service

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