FROM node:16-alpine
WORKDIR /app
COPY ./adminpanel-ui .
RUN npm ci
RUN npm run build
ENV NODE_ENV "production"
ENV REACT_APP_BACKEND="http://server:8000"
ENV REACT_APP_MINIO="http://server:8000/static/"
EXPOSE 3000
CMD [ "npx", "serve", "build" ]