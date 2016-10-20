FROM scratch
MAINTAINER Martin Marcher <martin@marcher.name>

COPY httpHostname /httpHostname
ENV LISTEN_PORT 8080
EXPOSE 8080
CMD [ "/httpHostname" ]
