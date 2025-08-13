#!/bin/bash
set -e

# Install required packages
#apk add --no-cache curl jq
#apk add --no-cache jq

# Wait for RabbitMQ to be ready
echo "Waiting for RabbitMQ to be ready..."
until curl -s -u admin:admin http://rabbitmq:15672/api/overview > /dev/null 2>&1; do
    echo "RabbitMQ not ready yet, waiting..."
    sleep 5
done

echo "RabbitMQ is ready! Starting subscriber..."

# Wait a bit more for the publisher to set up the queue
sleep 10

echo "Starting to consume messages..."

while true; do
    # Try to get a message from the queue
    response=$(curl -s -u admin:admin -X POST \
        -H "Content-Type: application/json" \
        -d '{"count":1,"ackmode":"ack_requeue_false","encoding":"auto","truncate":50000}' \
        http://rabbitmq:15672/api/queues/%2F/demo_queue/get)
    
    # Check if we got a message
    if echo "$response" | jq -e '.[0]' > /dev/null 2>&1; then
        # Extract the message payload
        message=$(echo "$response" | jq -r '.[0].payload')
        timestamp=$(date '+%Y-%m-%d %H:%M:%S')
        echo "[${timestamp}] Received: ${message}"
    else
        echo "No messages available, waiting..."
    fi
    
    sleep 3
done

