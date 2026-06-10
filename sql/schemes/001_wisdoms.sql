-- +goose Up
create table wisdoms (
    id integer primary key,
    data text,
    author text not null
);
-- +goose Down
drop table wisdoms;