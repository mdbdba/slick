FROM alpine:latest

LABEL maintainer="mdbntl@gmail.com"

WORKDIR /opt/slick
ADD ./cmd /opt/slick

RUN chmod -R a+r /opt/slick \
 && find /opt/slick -type d -print0 | xargs -0 chmod a+rx \
 && chmod 775 /opt/slick/slick \
 && chown -R 1337:1337 /opt/slick

USER 1337
CMD sh -c './slick'
