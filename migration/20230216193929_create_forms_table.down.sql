CREATE TABLE forms(
    id serial,
    attributes json null,
    rules json null,
    sort numeric null default 0 
)