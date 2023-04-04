FROM node:alpine as base
WORKDIR /frontend
COPY . .
RUN npm install


FROM base as dev
EXPOSE 80
CMD [ "npm", "run", "dev" ]

FROM base as builder
RUN npm run build

FROM node:alpine
WORKDIR /frontend
COPY --from=builder /frontend/build ./build
EXPOSE 80
CMD [ "npm", "run", "start" ]