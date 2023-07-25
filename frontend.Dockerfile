FROM node:16-alpine as build-stage
WORKDIR /app
COPY adminpanel-ui/package.json adminpanel-ui/package-lock.json ./
RUN npm ci
COPY ./adminpanel-ui .
ENV NODE_ENV="production"
ENV REACT_APP_BACKEND="http://localhost:8000"
ENV REACT_APP_MINIO="http://localhost:8000/static/"
RUN npm run build

FROM nginx:latest
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build-stage /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]

#FROM node:16-alpine as builder-stage
#WORKDIR /app
#COPY ./adminpanel-ui .
#RUN npm ci
#ENV NODE_ENV="production"
#ENV REACT_APP_BACKEND="http://188.225.44.3:8000"
#ENV REACT_APP_MINIO="http://188.225.44.3:8000/static/"
#RUN npm run build
#EXPOSE 3000
#CMD [ "npx", "serve", "build" ]