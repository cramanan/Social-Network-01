ARG NEXT_PUBLIC_API_URL=localhost:3001

FROM node:23-alpine AS deps
COPY package*.json ./
RUN npm ci

FROM node:23-alpine AS builder
COPY . .
COPY --from=deps ./node_modules ./node_modules
RUN npm run build

FROM node:23-alpine AS runner
COPY --from=builder /next.config.ts ./
COPY --from=builder /public ./public
COPY --from=builder /node_modules ./node_modules
COPY --from=builder /.next ./.next

EXPOSE 3000

CMD ["node_modules/.bin/next", "start"]