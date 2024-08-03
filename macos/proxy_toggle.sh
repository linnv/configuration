#!/bin/bash

# Configuration
# NETWORK_SERVICE="Wi-Fi"  # Change this to your network service (e.g., "Ethernet" or "Wi-Fi")
NETWORK_SERVICE="Ethernet"  # Change this to your network service (e.g., "Ethernet" or "Wi-Fi")
PROXY_HOST="127.0.0.1"  # Change this to your SOCKS proxy host
PROXY_PORT="8004"  # Change this to your SOCKS proxy port
# PROXY_PORT="8005"  # Change this to your SOCKS proxy port
BYPASS_FILE="$HOME/git/configuration/macos/.socks_proxy_bypass_domains.txt"  # File to store bypass domains

# Function to display usage
usage() {
    echo "Usage: $0 [on|off|status|add-bypass|remove-bypass|list-bypass]"
    echo "  on: Turn SOCKS proxy on"
    echo "  off: Turn SOCKS proxy off"
    echo "  status: Show current SOCKS proxy status"
    echo "  add-bypass <domain>: Add a domain to bypass list"
    echo "  remove-bypass <domain>: Remove a domain from bypass list"
    echo "  list-bypass: List all bypass domains"
    exit 1
}

# Function to turn SOCKS proxy on
proxy_on() {
    echo "Turning SOCKS proxy on..."
    networksetup -setsocksfirewallproxy "$NETWORK_SERVICE" "$PROXY_HOST" "$PROXY_PORT"
    apply_bypass_rules
    echo "SOCKS proxy is now ON."
}

# Function to turn SOCKS proxy off
proxy_off() {
    echo "Turning SOCKS proxy off..."
    networksetup -setsocksfirewallproxystate "$NETWORK_SERVICE" off
    echo "SOCKS proxy is now OFF."
}

# Function to show SOCKS proxy status
proxy_status() {
    echo "SOCKS Proxy status:"
    networksetup -getsocksfirewallproxy "$NETWORK_SERVICE"
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
