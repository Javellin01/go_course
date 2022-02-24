FROM mongo:latest
ENV MONGO_INITDB_ROOT_USERNAME=root
ENV MONGO_INITDB_ROOT_PASSWORD=mongodb
ENV MONGO_INITDB_DATABASE=mongodb
COPY deployments/docker/mongo/mongo-init.js /docker-entrypoint-initdb.d/mongo-init.js
EXPOSE 27017
