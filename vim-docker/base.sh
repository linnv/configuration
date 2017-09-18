#!/usr/bin/env bash

apt-get -y --no-install-recommends install \
  aptitude \
  curl \
  cmake \
  bash-completion \
  build-essential \
  python-dev \
  python-pip \
  software-properties-common \
  clang \
  coreutils \
  gcc \
  g++ \
  git-core \
  golang \
  mercurial \
  vim


add-apt-repository -y  ppa:neovim-ppa/stable
apt-get -y update
apt-get -y install neovim ctags 
pip install --upgrade pip
pip install setuptools
pip install distribute
pip install neovim

curl -fLo ~/.local/share/nvim/site/autoload/plug.vim --create-dirs \
    https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim

mkdir -p ~/.config/nvim
# TODO more check
mv /init.vim ~/.config/nvim 

mkdir  -p ~/go
export GOPATH=~/go
export GOBIN=~/go/bin
GOPATH="$GOPATH" go get -u -t -v github.com/golang/tools

mv "$GOPATH"/src/github.com "$GOPATH"/src/golang.org
mv "$GOPATH"/src/golang.org/golang "$GOPATH"/src/golang.org/x
# edit abs path of cmd:  <18-09-17, yourname> #
SCRIPT_DIR="$GOPATH"/src/golang.org/x/tools/cmd
for i in "$SCRIPT_DIR"/*; do
	echo "$i"
	cd "$i"
	GOPATH="$GOPATH" go install
done

cat >> ~/.bashrc<<-'endedit'

export GOPATH=/root/go
export GOBIN=$GOPATH/bin
export PATH=$GOBIN:$PATH
export GOOS=linux

alias nv=nvim
alias gst="git status"
alias gco="git checkout"
alias gb="git branch"
alias giff="git diff"
endedit

