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

CREATE TYPE booking_status AS ENUM ('reserved', 'checking', 'checkout', 'canceled');

create table booking (
    id serial primary key,
    customer_id int NOT NULL,
    room_id int NOT NULL,
    start_datetime timestamp NOT NULL,
    end_datetime timestamp NOT NULL,
    status booking_status NOT NULL,
    parking bool NOT NULL,
    foreign key (customer_id) REFERENCES customer (id),
    foreign key (room_id) REFERENCES room (id)
);

create table checkin (
    id serial primary key,
    booking_id int NOT NULL,
    checking_datetime timestamp NOT NULL,
    checkout_datetime timestamp,
    foreign key (booking_id) REFERENCES booking (id)
);

create table bill (
    id serial primary key,
    booking_id int NOT NULL,
    total_value numeric(5, 2),
    extra_hour bool,
    foreign key (booking_id) REFERENCES booking (id)
);

CREATE TYPE payments AS ENUM ('credit card', 'cash', 'pix');

create table payment (
    id serial primary key,
    bill_id int NOT NULL,
    value numeric(5, 3),
    type_payment payments,
    installments int,
    foreign key (bill_id) REFERENCES bill (id)
);

INSERT INTO customer ("name","document",phone_number,email) VALUES
	 ('Vanessa','44154875485','11966698722','vanessa.nunes@hotmail.com');

INSERT INTO room (room_number,description) VALUES
	 (1,'Quarto Azul'),
	 (2,'quarto Rosa');

INSERT INTO booking (customer_id,room_id,start_datetime,end_datetime,status,parking) VALUES
	 (1,2,'2023-05-17 19:17:40.26867','2023-05-18 14:00:00','checking',true);

INSERT INTO checkin (booking_id,checking_datetime,checkout_datetime) VALUES
	 (1,'2023-05-17 16:01:00','2023-05-18 16:01:00');



