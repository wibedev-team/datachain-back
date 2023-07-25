FROM node:16-alpine as build-stage
WORKDIR /app
COPY adminpanel-ui/package.json adminpanel-ui/package-lock.json ./
RUN npm ci
COPY ./adminpanel-ui .
ENV NODE_ENV="production"
RUN npm run build

FROM nginx:latest
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build-stage /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
