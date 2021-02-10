# build container
FROM node:lts-alpine as builder
WORKDIR /app
COPY . .
RUN yarn install && yarn build

# app container
FROM node:lts-alpine
WORKDIR /app

# this is unavoidable because of nuxt, unless i want to not include SEO
COPY --from=builder /app/node_modules node_modules/
COPY --from=builder /app/.nuxt .nuxt/
COPY package.json .
COPY nuxt.config.ts .
COPY tsconfig.json .
COPY static .

EXPOSE 3000
ENV HOST=0.0.0.0
ENV TEST=TICKLE
CMD [ "yarn", "start" ]
