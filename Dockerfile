FROM postgres:10.3-alpine
COPY data.sql /docker-entrypoint-initdb.d/
EXPOSE 5432