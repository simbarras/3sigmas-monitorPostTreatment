FROM node:20-alpine as builder
WORKDIR /app
COPY package.json .

RUN npm install --save-dev
COPY . .
RUN npm run build

FROM nginx:stable-alpine as production
COPY --from=builder /app/dist /usr/share/nginx/html
EXPOSE 80
ENTRYPOINT ["nginx","-g","daemon off;"]