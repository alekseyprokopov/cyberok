CREATE TABLE whois
(
    id     bigserial not null primary key,
    domain varchar   not null unique,
    info   varchar   not null
);