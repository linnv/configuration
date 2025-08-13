FROM alpine:latest

# Install curl
RUN http_proxy=http://192.168.0.6:9015 https_proxy=http://192.168.0.6:9015 apk --no-cache add curl jq

# Set default command
CMD ["sh"]

