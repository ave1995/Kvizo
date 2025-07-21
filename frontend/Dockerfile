# Stage 1: Build the React app
FROM node:24-alpine AS build

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .
RUN npm run build

# Stage 2: Serve with a simple static server (e.g., nginx)
FROM nginx:alpine

# Remove the default nginx config
RUN rm /etc/nginx/conf.d/default.conf

# Copy your nginx.conf into the container
COPY docker/nginx.conf /etc/nginx/nginx.conf

# Copy built frontend assets to nginx serving directory
COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
