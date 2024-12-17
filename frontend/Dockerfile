FROM node:20-slim

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm install

COPY . .

RUN npm run build
EXPOSE 5178

CMD ["npx", "vite", "preview", "--host", "0.0.0.0", "--port", "5178"]
