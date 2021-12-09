#FROM node:alpine as build
#
#WORKDIR /app
#
#COPY ./package.json ./
#
#RUN npm install
#
#COPY . .
#
#RUN npm run build
#
#
#
##FROM node:alpine
#
##WORKDIR /app
#
##COPY --from=build /app/build .
#
#RUN npm install -g serve
#
#CMD ["serve", "-s", "build"]

FROM node:alpine as build
WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH
COPY package.json /app/package.json
RUN npm install
COPY . /app
RUN npm run build

# production environment
FROM node:alpine
COPY --from=build /app/build /app/build
WORKDIR /app

ENV PORT 3000
EXPOSE 3000

RUN npm install -g serve
CMD ["serve", "-s", "build"]