FROM nginx:latest

WORKDIR /dist

COPY ./dist /dist
COPY nginx.conf.template /
COPY . /home

CMD envsubst < /nginx.conf.template > /etc/nginx/nginx.conf \
	&& nginx -g 'daemon off;'
