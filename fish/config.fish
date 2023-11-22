if status is-interactive
    # Commands to run in interactive sessions can go here
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


# Set personal aliases, overriding those provided by oh-my-zsh libs,
# # users are encouraged to define aliases within the ZSH_CUSTOM folder.
# # For a full list of active aliases, run `alias`.
# # Example aliases
# # alias zshconfig="mate ~/.zshrc"
# # alias ohmyzsh="mate ~/.oh-my-zsh"
# # alias ohmyzsh="mate ~/.oh-my-zsh"
# alias cdwindownload='cd /mnt/c/Users/jialinwu/Downloads'
# alias cdwintmp='cd /mnt/c/Users/jialinwu/tmpwin'
# alias cdds='cd /mnt/c/Users/jialinwu/Sync'
# alias gitp="git pull origin"
# alias git merge="git merge --no--ff"
alias gpo='git push origin'
# # git config --global alias.dsf '!f() { [ -z "$GIT_PREFIX" ] || cd "$GIT_PREFIX" '\
	# # # alias giff="git diff --color | diff-so-fancy | less"
alias giff="git dsf"
# alias python="/usr/bin/python3"
alias giff="git diff"
alias gst="git status"
alias cdgo="cd ~/go/src"
# alias klogs="kubectl logs -f "
# alias tmux="/snap/bin/tmux-non-dead.tmux")
alias gcmsg="git commit -m"
alias gco="git checkout"
alias gb="git branch"

alias nvim "no_proxy=127.0.0.1,localhost,gitlab.qnzsai.com http_proxy=http://100.108.162.74:8015 https_proxy=http://100.108.162.74:8015 command nvim"
alias nv='nvim'
alias vim='nvim'

# alias vim="/opt/homebrew/bin/vim"
alias glog="git log --oneline --decorate --graph"
alias gloga="git log --oneline --decorate --graph --all"

alias grep="/opt/homebrew/Cellar/grep/3.11/bin/ggrep"
