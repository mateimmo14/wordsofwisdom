-- +goose Up
create table wisdoms (
    id integer generated always as identity,
    data text,
    author text not null
);
-- +goose Down
drop table wisdoms;