create table if not exists burgers(
id bigserial primary key ,
name text not null,
price int,
remove bool
);


drop table burgers

create table if not exists burgers(
                                      id bigserial primary key ,
                                      name text not null,
                                      price int,
                                      remove bool
);