FROM postgres:10.6

ENV POSTGRES_USER transaction-service-app
ENV POSTGRES_DB transaction-service

# migrations
COPY internal/migration/sql/* /docker-entrypoint-initdb.d/