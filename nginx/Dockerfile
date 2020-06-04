FROM node:10-alpine as builder

# copy the package.json to install dependencies
#RUN mkdir skillbased-ui

#COPY frontend/package.json frontend/package-lock.json ./skillbased-ui
WORKDIR skillbased-ui
COPY frontend/ .

#RUN npm install && mkdir /skillbased-ui && mv ./node_modules ./skillbased-ui
RUN npm install
#COPY . ./skillbased-ui

#WORKDIR skillbased-ui

#COPY . .

# Build the project and copy the files
RUN npm run build


FROM nginx:alpine

#!/bin/sh

COPY nginx/nginx.conf /etc/nginx/nginx.conf

## Remove default nginx index page
RUN rm -rf /usr/share/nginx/html/*

# Copy from the stahg 1
COPY --from=builder /skillbased-ui/build /usr/share/nginx/html

EXPOSE 3000 80

ENTRYPOINT ["nginx", "-g", "daemon off;"]