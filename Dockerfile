FROM debian:stable-slim
RUN mkdir -p /data

ADD ./dist/static /data/static
ADD ./dist/templates /data/templates
ADD ./dist/config.json /data/config.json
ADD ./dist/snowfall.linux /data/snowfall.linux

WORKDIR /data
ENV GIN_MODE=release
ENV PORT=20000

HEALTHCHECK --interval=8s --timeout=1s \
    CMD curl -f http://127.0.0.1:20000/ || exit 1

EXPOSE 20000
CMD ["./snowfall.linux"]
