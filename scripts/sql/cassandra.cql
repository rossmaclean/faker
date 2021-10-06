create table address
(
    address_id      varchar,
    building_number varchar,
    street_address  varchar,
    city            varchar,
    country         varchar,
    post_code       varchar,
    primary key (address_id)
);

create table person
(
    person_id          varchar,
    first_name         varchar,
    last_name          varchar,
    date_of_birth      varchar,
    address_id         varchar,
    phone_number       varchar,
    email_address      varchar,
    credit_card_number varchar,
    primary key (person_id)
);
