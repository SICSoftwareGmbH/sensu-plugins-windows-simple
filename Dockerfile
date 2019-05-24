# Build image
FROM golang:1.12-alpine

# Maintainer
MAINTAINER Florian Schwab <florian.schwab@sic.software>

# Upgrade system
RUN apk --no-cache --no-progress --update upgrade

# Install os dependencies
RUN apk --no-cache --no-progress --update add bash build-base curl git ca-certificates

# Set the working directory
WORKDIR /go/src/sensu-plugins-windows-simple

# Default command
CMD ["bash"]
