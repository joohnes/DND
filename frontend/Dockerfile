FROM node:latest AS frontend
WORKDIR /app/frontend
COPY ./frontend/package*.json ./
RUN npm install
COPY ./frontend .
RUN npm run build
EXPOSE 5178
CMD ["node", "build/index.js"]
