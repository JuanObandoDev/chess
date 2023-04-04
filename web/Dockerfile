FROM node:alpine as base
WORKDIR /web
COPY . .
RUN npm install
ENV PORT=80
EXPOSE 80


FROM base as dev
CMD [ "npm", "run", "dev" ]

FROM base
RUN npm run build
CMD [ "npm", "run", "start" ]