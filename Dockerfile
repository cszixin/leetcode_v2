FROM golang:1.17.7-buster
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone
RUN echo 'alias ls="ls --color=auto"' >>  ~/.bashrc