FROM node:14.15-alpine
LABEL maintainer="Julio Cesar <julio@blackdevs.com.br>"

WORKDIR /usr/src/app

COPY package.json yarn.lock ./
RUN yarn install --production=true
COPY . .

CMD ["node", "index.js"]
