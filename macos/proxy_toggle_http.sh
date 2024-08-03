#!/bin/bash

# Configuration
NETWORK_SERVICE="Ethernet"  # Change this to your network service (e.g., "Ethernet" or "Wi-Fi")
PROXY_HOST="127.0.0.1"  # Change this to your HTTP proxy host
PROXY_PORT="8014"  # Change this to your HTTP proxy port
BYPASS_FILE="$HOME/git/configuration/macos/.http_proxy_bypass_domains.txt"  # File to store bypass domains

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
    networksetup -setwebproxy "$NETWORK_SERVICE" "$PROXY_HOST" "$PROXY_PORT"
    networksetup -setsecurewebproxy "$NETWORK_SERVICE" "$PROXY_HOST" "$PROXY_PORT"
    apply_bypass_rules
    echo "HTTP proxy is now ON."
}

# Function to turn HTTP proxy off
proxy_off() {
    echo "Turning HTTP proxy off..."
    networksetup -setwebproxystate "$NETWORK_SERVICE" off
    networksetup -setsecurewebproxystate "$NETWORK_SERVICE" off
    echo "HTTP proxy is now OFF."
}

# Function to show HTTP proxy status
proxy_status() {
    echo "HTTP Proxy status:"
    networksetup -getwebproxy "$NETWORK_SERVICE"
    networksetup -getsecurewebproxy "$NETWORK_SERVICE"
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
        networksetup -setproxybypassdomains "$NETWORK_SERVICE" $BYPASS_DOMAINS
        echo "Bypass rules applied."
    else
        networksetup -setproxybypassdomains "$NETWORK_SERVICE" "Empty"
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
