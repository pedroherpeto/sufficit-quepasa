FROM golang:1.17-stretch

WORKDIR /app

COPY docker-entrypoint.sh /usr/local/bin/

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENTRYPOINT ["docker-entrypoint.sh"]