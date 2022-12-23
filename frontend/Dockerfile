FROM node:18.12-alpine as build

WORKDIR /node/frontend/src

COPY . .

RUN npm install
# RUN npm ci --only=production

RUN npm run build

FROM node:18.12-alpine

ENV NODE_ENV=production

WORKDIR /node/frontend/dist

COPY --from=build /node/frontend/dist /node/frontend/dist

COPY --from=build /node/frontend/src/package*.json /node/frontend/dist/

ENTRYPOINT [ "node", "/node/frontend/dist/index.js" ]
# trigger pipeline 10