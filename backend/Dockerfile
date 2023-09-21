FROM golang:1.19-alpine

RUN apk update && \
    apk add --no-cache git bash curl

WORKDIR /admin

COPY . .

RUN chmod +x /admin/update.sh
RUN chmod +x /admin/run.sh
RUN sh /admin/update.sh

EXPOSE 9123
EXPOSE 9991


CMD ["sh","/admin/run.sh"]