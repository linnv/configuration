#!/bin/bash

setup_macos_defaults() {
  echo "==========================================================="
  echo "               Setting up macOS Defaults...                "
  echo "-----------------------------------------------------------"

  echo "Setting Dock to auto-hide..."
  defaults write com.apple.dock "autohide" -bool "true"

  echo "Setting faster key repeat rate for Vim users..."
  # Set a faster key repeat rate (lower value = faster)
  defaults write NSGlobalDomain KeyRepeat -int 2
  # Set a shorter delay until repeat starts (lower value = shorter delay)
  defaults write NSGlobalDomain InitialKeyRepeat -int 15

  #Show full website URL in Safari
  defaults write com.apple.Safari "ShowFullURLInSmartSearchField" -bool "true" && killall Safari

  # Show hidden files in Finder
  defaults write com.apple.finder "AppleShowAllFiles" -bool "true"

  # Disable the warning when changing a file extension
  defaults write com.apple.finder FXEnableExtensionChangeWarning -bool false

  # Show time in menu bar
  defaults write com.apple.menuextra.clock "DateFormat" -string "\"EEE HH:mm:ss\""
  defaults write com.apple.menuextra.clock ShowSeconds -bool false

  echo "Restarting affected applications..."
  killall Dock
  killall Finder

  echo "macOS defaults have been updated!"
}

setup_macos_defaults
