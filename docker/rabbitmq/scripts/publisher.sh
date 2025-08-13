#!/bin/bash
set -e

# Install required packages
#apk add --no-cache curl

# Wait for RabbitMQ to be ready
echo "Waiting for RabbitMQ to be ready..."
until curl -s -u admin:admin http://rabbitmq:15672/api/overview > /dev/null 2>&1; do
    echo "RabbitMQ not ready yet, waiting..."
    sleep 5
done

echo "RabbitMQ is ready! Setting up exchange and queue..."

# Create exchange
curl -u admin:admin -X PUT \
    -H "Content-Type: application/json" \
    -d '{"type":"direct","auto_delete":false,"durable":true,"arguments":{}}' \
    http://rabbitmq:15672/api/exchanges/%2F/demo_exchange

# Create queue
curl -u admin:admin -X PUT \
    -H "Content-Type: application/json" \
    -d '{"auto_delete":false,"durable":true,"arguments":{}}' \
    http://rabbitmq:15672/api/queues/%2F/demo_queue

# Bind queue to exchange
curl -u admin:admin -X POST \
    -H "Content-Type: application/json" \
    -d '{"routing_key":"demo.message","arguments":{}}' \
    http://rabbitmq:15672/api/bindings/%2F/e/demo_exchange/q/demo_queue

echo "Exchange and queue setup complete!"

# Start publishing messages
echo "Starting to publish messages..."
counter=1

while true; do
    timestamp=$(date '+%Y-%m-%d %H:%M:%S')
    message="Hello from publisher! Message #${counter} at ${timestamp}"
    
    # Publish message
    curl -u admin:admin -X POST \
        -H "Content-Type: application/json" \
        -d "{\"properties\":{},\"routing_key\":\"demo.message\",\"payload\":\"${message}\",\"payload_encoding\":\"string\"}" \
        http://rabbitmq:15672/api/exchanges/%2F/demo_exchange/publish
    
    echo "Published: ${message}"
    
    counter=$((counter + 1))
    sleep 5
done
