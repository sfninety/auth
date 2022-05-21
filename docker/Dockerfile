FROM golang:alpine AS builder

# Install git
RUN apk update && apk add --no-cache git

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

# allow private repo pull
RUN git config --global url."https://${{ secrets.AUTH_TOKEN }}:x-oauth-basic@github.com/".insteadOf "https://github.com/"

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

# Exec as an unprivileged user
USER builduser:builduser

EXPOSE 8000

ENTRYPOINT [ "/go/bin/service" ]