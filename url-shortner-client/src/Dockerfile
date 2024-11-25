FROM node:22-alpine

WORKDIR /url-shorter-client

COPY package.json .
RUN npm install

COPY . .

EXPOSE 4000

CMD ["npm", "start"]
