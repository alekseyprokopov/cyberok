CREATE TABLE IF NOT EXISTS fqdn
(
    id   bigserial not null primary key,
    name varchar   not null unique
);

