FROM postgres:alpine
ADD scripts/1_init.up.sql /docker-entrypoint-initdb.d
RUN chmod a+r /docker-entrypoint-initdb.d/*
EXPOSE 5432