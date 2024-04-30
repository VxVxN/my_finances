FROM node:18-alpine AS deps
RUN apk add --no-cache libc6-compat
WORKDIR /web-app

COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile

FROM node:18-alpine AS builder
WORKDIR /web-app
COPY --from=deps /web-app/node_modules ./node_modules
COPY . .

ENV NEXT_TELEMETRY_DISABLED 1

RUN npm run build

FROM node:18-alpine AS runner
WORKDIR /web-app

ENV NODE_ENV production 

COPY --from=builder --chown=nextjs:nodejs /web-app/.next ./.next
COPY --from=builder /web-app/node_modules ./node_modules
COPY --from=builder /web-app/package.json ./package.json

EXPOSE 3000

ENV PORT 3000

CMD ["npm", "start"]
