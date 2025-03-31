function run_for_aarch64arm64
    echo "Executing function for arm64 (aarch64)"
    /opt/homebrew/bin/brew shellenv | source
    set -x CXX /opt/homebrew/opt/llvm/bin/clang++
    alias grep="/opt/homebrew/Cellar/grep/3.11/bin/ggrep"
    # Add architecture-specific commands here
end

function run_for_amd64
    echo "Executing function for x86_64 (amd64)"
    # Add architecture-specific commands here
end

# Detect architecture and execute corresponding function
switch (uname -m)
    case aarch64
    case arm64
        run_for_aarch64arm64
    case x86_64
        run_for_amd64
    case '*'
        echo "Unsupported architecture: $(uname -m)"
end

if status is-interactive
    # Commands to run in interactive sessions can go here
# Prevent atuin from applying its default keybindings
    atuin init fish | source
    starship init fish | source
end

#use ltime to show last cmd time-elapsed
function ltime --description 'Print command duration in seconds for last command'
	    echo (echo 'scale=3; ' $CMD_DURATION ' / 1000' | bc)"s"
end

function printtime --on-event fish_postexec
	    # set duration (echo 'scale=3; ' $CMD_DURATION ' / 1000' | bc)"s"
	 set duration (echo "$CMD_DURATION 1000" | awk '{printf "%.3fs", $1 / $2}')
        echo -e "\nelapsed $duration"
end

alias gpo='git push origin'
alias giff="git dsf"
alias giff="git diff"
alias gst="git status"
alias cdgo="cd ~/go/src"
alias gcmsg="git commit -m"
alias gco="git checkout"
alias gb="git branch"

alias glog="git log --oneline --decorate --graph"
alias gloga="git log --oneline --decorate --graph --all"

function fzf_switch_window
	    tmux list-windows -a -F "#S:#I-#W" | fzf-tmux | cut -d "-" -f 1 | xargs tmux switch-client -t
end

alias nvim "no_proxy=127.0.0.1,localhost,gitlab.qnzsai.com http_proxy=http://codeium.nat:8015 https_proxy=http://codeium.nat:8015 command nvim"
alias nv='vim'
alias vim "no_proxy=127.0.0.1,localhost,gitlab.qnzsai.com http_proxy=http://codeium.nat:8015 https_proxy=http://codeium.nat:8015 command vim"
alias vi "no_proxy=127.0.0.1,localhost,gitlab.qnzsai.com http_proxy=http://codeium.nat:8015 https_proxy=http://codeium.nat:8015 command vim"

alias ssh="TERM=xterm-256color command ssh"

