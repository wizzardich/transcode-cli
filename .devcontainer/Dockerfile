# Courtesy of https://morioh.com/p/de68c84561e5

ARG GO_VERSION=1.18
ARG DEBIAN_VERSION=bullseye

FROM golang:${GO_VERSION}-${DEBIAN_VERSION}

ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=1000

# Install packages and Go language server

RUN apk add -q --update --progress --no-cache git sudo openssh-client zsh gnupg

RUN go get -u -v golang.org/x/tools/cmd/gopls 2>&1

# Setup user

RUN useradd -s /bin/zsh -D -u $USER_UID -g $USER_GID $USERNAME && \
    mkdir -p /etc/sudoers.d && \
    echo $USERNAME ALL=(root) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME && \
    chmod 0440 /etc/sudoers.d/$USERNAME

# Setup shell

USER $USERNAME

RUN sh -c “$(wget -O- https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh)” “” --unattended &> /dev/null

ENV ENV=“/home/$USERNAME/.ashrc” \
    ZSH=/home/$USERNAME/.oh-my-zsh \ 
    EDITOR=vi \
    LANG=en_US.UTF-8

RUN printf ‘ZSH_THEME=“robbyrussell”\nENABLE_CORRECTION=“false”\nplugins=(git copyfile extract colorize dotenv encode64 golang vscode)\nsource $ZSH/oh-my-zsh.sh’ > “/home/$USERNAME/.zshrc”

USER root