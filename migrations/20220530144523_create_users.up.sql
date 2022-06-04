CREATE TABLE users (
    id bigserial not null primary key,
    email varchar not null unique,
    encrypted_password varchar not null
);

CREATE TABLE competitions (
    id bigserial not null primary key,
    name varchar not null,
    dt timestamp not null,
    age_category varchar not null,
    weapon_type varchar not null,
    is_team int not null check (is_team < 2 and is_team > -1),
    status varchar,
    sex varchar not null check (sex = 'female' or sex = 'male'),
    type varchar
);
