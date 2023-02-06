CREATE TABLE IF NOT EXISTS users
    (
    id serial not null unique,
    firstname varchar(255) ,
    lastname varchar(255) ,
    username varchar(255) not null unique,
    email varchar(255) not null,
    password_hash varchar(255) not null,
    status varchar(255) ,
    balance float4
    );



CREATE TABLE IF NOT EXISTS currencies
(
    id serial not null unique,
    name varchar(255) not null unique,
    value float4 not null
);



CREATE TABLE IF NOT EXISTS basket
(
    id serial not null unique,
    user_id int references users(id)  not null,
    TransactionType varchar(255) not null,
    currency_id int references currencies(id)  not null,
    Currency varchar(255) references currencies(name)  not null,
    Value float4 ,
    amount float4 ,
    Status varchar(255) not null
);

INSERT INTO currencies (name, value)  values ('usdt', 1);


