############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git build-base
WORKDIR $GOPATH/src/mypackage/myapp/
COPY .github/workflows .
# Build the binary.
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux GIN_MODE=release go build -ldflags="-w -s" -o /go/bin/postTreatment cmd/postTreatment/main.go

############################
# STEP 2 build a small image
############################

FROM alpine AS postTreatment
# Copy our static executable.
COPY --from=builder /go/bin/postTreatment /go/bin/app
ENTRYPOINT ["/go/bin/app"]

