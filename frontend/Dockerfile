FROM nginx:1.15.12-alpine

COPY dist/front /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

RUN apk update && apk add git

WORKDIR /usr/share/nginx/html/src/js/lib

EXPOSE 80
