FROM node:12-alpine

MAINTAINER balamaru

ADD app.js app.js

CMD ["node", "app.js"]
