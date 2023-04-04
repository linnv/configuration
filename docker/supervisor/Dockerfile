FROM ubuntu:20.04

RUN apt-get update && \
    apt-get install -y vim git wget less rsync tmux && \
    apt-get install -y supervisor && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN sed -i 's@/etc/supervisor/conf.d/\*\.conf@/etc/supervisor/conf.d/*.conf /etc/supervisor/conf.d/*.conf /data/supervisor/conf.d/*.conf@g' /etc/supervisor/supervisord.conf

CMD ["/usr/bin/supervisord", "-n"]


