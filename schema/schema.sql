
create table if not exists burgers(
                                      id bigserial primary key ,
                                      name text not null,
                                      price int,
                              removed bool
);