-- +goose Up
-- +goose StatementBegin
CREATE TABLE segments
(
  id serial unique not null ,
  name varchar(255) not null unique ,
  description varchar(255)
);

CREATE TABLE users
(
    id serial unique not null references segments(id),
    user_segments varchar(255)

);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE segments;
DROP TABLE users;
-- +goose StatementEnd
