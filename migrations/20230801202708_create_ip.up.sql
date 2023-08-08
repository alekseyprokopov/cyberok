CREATE TABLE IF NOT EXISTS ip
(
    id      bigserial not null primary key,
    fqdn_id int references fqdn(id),
    ip      varchar   not null
);

