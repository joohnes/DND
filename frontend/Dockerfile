FROM node:latest AS frontend
WORKDIR /app/frontend
COPY ./frontend/package*.json ./
RUN npm install
COPY ./frontend .
EXPOSE 5178
CMD ["npm", "run", "dev", "--", "--host", "--port", "5178"]
