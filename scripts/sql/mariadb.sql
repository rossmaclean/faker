create table address
(
    address_id      varchar(255) not null,
    building_number varchar(255) not null,
    street_address  varchar(255) not null,
    city            varchar(255) not null,
    country         varchar(255) not null,
    post_code       varchar(255) not null,
    constraint address_pk
        primary key (address_id)
);

create table person
(
    person_id          varchar(255) not null,
    first_name         varchar(255) not null,
    last_name          varchar(255) not null,
    date_of_birth      varchar(255) not null,
    address_id         varchar(255) not null,
    phone_number       varchar(255) not null,
    email_address      varchar(255) not null,
    credit_card_number varchar(255) not null,
    constraint person_pk
        primary key (person_id),
    constraint person_address_address_id_fk
        foreign key (address_id) references address (address_id)
);