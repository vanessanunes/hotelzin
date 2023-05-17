create table customer (
    id serial primary key,
    name varchar(50) NOT NULL,
    document varchar(11) UNIQUE NOT NULL,
    phone_number varchar(11) NOT NULL,
    email varchar(50)
);

create table room (
    id serial primary key,
    room_number int UNIQUE NOT NULL,
    description varchar(50)
);

create table booking (
    id serial primary key,
    customer_id int NOT NULL,
    room_id int NOT NULL,
    started_datetime date,
    finished_datetime date,
    status boolean,
    foreign key (customer_id) REFERENCES customer (id),
    foreign key (room_id) REFERENCES room (id)
);

create table checkin (
    id serial primary key,
    booking_id int NOT NULL,
    checking_datetime date,
    checkout_datetime date,
    foreign key (booking_id) REFERENCES booking (id)
);

create table bill (
    id serial primary key,
    booking_id int NOT NULL,
    extra_hour numeric(5, 2),
    discount numeric(5, 2),
    foreign key (booking_id) REFERENCES booking (id)
);

create table payment (
    id serial primary key,
    bill_id int NOT NULL,
    value numeric(5, 3),
    type_payment varchar(10),
    installments int,
    foreign key (bill_id) REFERENCES bill (id)
);