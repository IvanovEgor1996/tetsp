create table if not exists balances
(
    "user" uuid not null
    constraint balance_pk
    primary key,
    balance integer default 0
);

alter table balances owner to CURRENT_USER;

create unique index if not exists balance_user_uindex
    on balances ("user");
