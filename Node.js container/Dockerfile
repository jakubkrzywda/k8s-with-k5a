FROM node:22-alpine

WORKDIR /app

COPY package*.json /app

RUN npm install

USER node

COPY --chown=node:node app.js /app

CMD ["node", "app.js"]
