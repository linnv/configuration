#!/bin/bash

# Configuration
NETWORK_SERVICES=("Wi-Fi" "Ethernet")  # Add your network services here
PROXY_HOST="127.0.0.1"  # Change this to your SOCKS proxy host
PROXY_PORT="1096"  # Change this to your SOCKS (scoks of singbox/gost not woking,don't know why) proxy port
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

# Function to turn SOCKS proxy on for all network services
proxy_on() {
    echo "Turning SOCKS proxy on for all network services..."
    for SERVICE in "${NETWORK_SERVICES[@]}"; do
        echo "Configuring $SERVICE..."
        networksetup -setsocksfirewallproxy "$SERVICE" "$PROXY_HOST" "$PROXY_PORT"
    done
    apply_bypass_rules
    echo "SOCKS proxy is now ON for all network services."
}

# Function to turn SOCKS proxy off for all network services
proxy_off() {
    echo "Turning SOCKS proxy off for all network services..."
    for SERVICE in "${NETWORK_SERVICES[@]}"; do
        echo "Configuring $SERVICE..."
        networksetup -setsocksfirewallproxystate "$SERVICE" off
    done
    echo "SOCKS proxy is now OFF for all network services."
}

# Function to show SOCKS proxy status for all network services
proxy_status() {
    echo "SOCKS Proxy status for all network services:"
    for SERVICE in "${NETWORK_SERVICES[@]}"; do
        echo "Status for $SERVICE:"
        networksetup -getsocksfirewallproxy "$SERVICE"
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

# Function to apply bypass rules to all network services
apply_bypass_rules() {
    if [ -f "$BYPASS_FILE" ]; then
        BYPASS_DOMAINS=$(tr '\n' ' ' < "$BYPASS_FILE")
        for SERVICE in "${NETWORK_SERVICES[@]}"; do
            echo "Applying bypass rules to $SERVICE..."
            networksetup -setproxybypassdomains "$SERVICE" $BYPASS_DOMAINS
        done
        echo "Bypass rules applied."
    else
        for SERVICE in "${NETWORK_SERVICES[@]}"; do
            echo "Clearing bypass rules for $SERVICE..."
            networksetup -setproxybypassdomains "$SERVICE" "Empty"
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
