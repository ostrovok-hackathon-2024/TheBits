FROM node:20 as build
RUN npm install -g @quasar/cli

WORKDIR /app
COPY . .
RUN npm install
EXPOSE 8080
CMD [ "quasar", "dev", "--host", "0.0.0.0", "--port", "8080", "--mode", "spa" ]
