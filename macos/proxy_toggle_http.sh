#!/bin/bash

# Configuration
NETWORK_SERVICES=("Wi-Fi" "Ethernet")  # Add or modify network services as needed
PROXY_HOST="127.0.0.1"  # Change this to your HTTP proxy host
# PROXY_PORT="9015"  # Change this to your HTTP proxy port
PROXY_PORT="1095"  # Change this to your HTTP proxy port
BYPASS_FILE="$HOME/git/configuration/macos/.socks_proxy_bypass_domains.txt"  # File to store bypass domains

# Function to display usage
usage() {
    echo "Usage: $0 [on|off|status|add-bypass|remove-bypass|list-bypass]"
    echo "  on: Turn HTTP proxy on"
    echo "  off: Turn HTTP proxy off"
    echo "  status: Show current HTTP proxy status"
    echo "  add-bypass <domain>: Add a domain to bypass list"
    echo "  remove-bypass <domain>: Remove a domain from bypass list"
    echo "  list-bypass: List all bypass domains"
    exit 1
}

# Function to turn HTTP proxy on
proxy_on() {
    echo "Turning HTTP proxy on..."
    for service in "${NETWORK_SERVICES[@]}"; do
        networksetup -setwebproxy "$service" "$PROXY_HOST" "$PROXY_PORT"
        networksetup -setsecurewebproxy "$service" "$PROXY_HOST" "$PROXY_PORT"
    done
    apply_bypass_rules
    echo "HTTP proxy is now ON for all configured network services."
}

# Function to turn HTTP proxy off
proxy_off() {
    echo "Turning HTTP proxy off..."
    for service in "${NETWORK_SERVICES[@]}"; do
        networksetup -setwebproxystate "$service" off
        networksetup -setsecurewebproxystate "$service" off
    done
    echo "HTTP proxy is now OFF for all configured network services."
}

# Function to show HTTP proxy status
proxy_status() {
    echo "HTTP Proxy status:"
    for service in "${NETWORK_SERVICES[@]}"; do
        echo "Service: $service"
        networksetup -getwebproxy "$service"
        networksetup -getsecurewebproxy "$service"
        echo "---"
    done
}

# Function to add a bypass domain
add_bypass() {
    if [ -z "$1" ]; then
        echo "Please specify a domain to add."
        exit 1
    fi
    echo "$1" >> "$BYPASS_FILE"
    echo "Added $1 to bypass list."
    apply_bypass_rules
}

# Function to remove a bypass domain
remove_bypass() {
    if [ -z "$1" ]; then
        echo "Please specify a domain to remove."
        exit 1
    fi
    sed -i '' "/^$1$/d" "$BYPASS_FILE"
    echo "Removed $1 from bypass list."
    apply_bypass_rules
}

# Function to list bypass domains
list_bypass() {
    if [ -f "$BYPASS_FILE" ]; then
        echo "Current bypass domains:"
        cat "$BYPASS_FILE"
    else
        echo "No bypass domains configured."
    fi
}

# Function to apply bypass rules
apply_bypass_rules() {
    if [ -f "$BYPASS_FILE" ]; then
        BYPASS_DOMAINS=$(tr '\n' ' ' < "$BYPASS_FILE")
        for service in "${NETWORK_SERVICES[@]}"; do
            networksetup -setproxybypassdomains "$service" $BYPASS_DOMAINS
        done
        echo "Bypass rules applied to all configured network services."
    else
        for service in "${NETWORK_SERVICES[@]}"; do
            networksetup -setproxybypassdomains "$service" "Empty"
        done
        echo "No bypass rules applied."
    fi
}

# Main script logic
case "$1" in
    on)
        proxy_on
        ;;
    off)
        proxy_off
        ;;
    status)
        proxy_status
        ;;
    add-bypass)
        add_bypass "$2"
        ;;
    remove-bypass)
        remove_bypass "$2"
        ;;
    list-bypass)
        list_bypass
        ;;
    *)
        usage
        ;;
esac

exit 0
