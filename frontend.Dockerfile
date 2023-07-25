FROM node:16-alpine as build-stage
WORKDIR /app
COPY adminpanel-ui/package.json adminpanel-ui/package-lock.json ./
RUN npm ci
COPY ./adminpanel-ui .
ENV NODE_ENV="production"
RUN npm run build

FROM nginx:latest

COPY nginx.conf /etc/nginx/conf.d/default.conf

COPY certs/admin.data-chainz.ru.crt /etc/nginx/certs/admin.data-chainz.ru.crt
COPY certs/admin.data-chainz.ru.key /etc/nginx/certs/admin.data-chainz.ru.key

COPY --from=build-stage /app/build /usr/share/nginx/html

EXPOSE 80
EXPOSE 443
CMD ["nginx", "-g", "daemon off;"]
