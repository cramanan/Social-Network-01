FROM node:22-bookworm

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY src ./src
COPY public ./public
COPY next.config.mjs .
COPY postcss.config.mjs .
COPY tailwind.config.ts .
COPY tsconfig.json .

EXPOSE 3000

RUN npm run build

CMD ["npm", "run", "start"]