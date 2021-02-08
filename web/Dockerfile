# build container
FROM node:lts-alpine as builder
WORKDIR /app
COPY . .
RUN yarn install \
	&& yarn build \
	&& yarn generate --fail-on-error

# app container
FROM nginx:alpine
COPY --from=builder /app/dist /usr/share/nginx/html
COPY --from=builder /app/nginx.conf /etc/nginx/conf.d/default.conf
