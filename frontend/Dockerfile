FROM node:20-alpine as builder
WORKDIR /app
COPY package.json .

RUN npm install --save-dev
COPY . .
RUN npm run build

FROM nginx:alpine as prdouction
RUN mkdir "/app"
COPY --from=builder /app/dist /app/dist
COPY nginx.conf /etc/nginx/nginx.conf
EXPOSE 80
ENTRYPOINT ["nginx","-g","daemon off;"]