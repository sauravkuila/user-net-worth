drop table if exists angel_one;
drop table if exists supported_broker;

create table angel_one(
    id serial primary key,
    symbol text not null,
    isin text not null,
    quantity int not null,
    price float not null,
    created_on timestamp default now(),
    updated_on timestamp default now()
);

create table idirect(
    id serial primary key,
    symbol text not null,
    isin text not null,
    quantity int not null,
    price float not null,
    created_on timestamp default now(),
    updated_on timestamp default now()
);

create table zerodha(
    id serial primary key,
    symbol text not null,
    isin text not null,
    quantity int not null,
    price float not null,
    created_on timestamp default now(),
    updated_on timestamp default now()
);

create table supported_broker(
    id serial primary key,
    broker_name text not null,
    created_on timestamp default now(),
    updated_on timestamp default now()
);

create table creds(
    id serial primary key,
    account text not null,
    totp_secret text,
    user_key text,
    pass_key text,
    app_api_key text,
    secret_key text,
    created_on timestamp default now(),
    updated_on timestamp default now()
)

create table broker_sync(
    id serial primary key,
    broker_id int not null,
    holdings_sync bool not null default FALSE,
    created_on timestamp default now(),
    updated_on timestamp default now(),
    foreign key (broker_id) references supported_broker(id)
);

create table holdings_updated

-- create function
CREATE  FUNCTION update_record()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_on = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- create trigger
CREATE TRIGGER sync_update_record
    BEFORE UPDATE
    ON
        supported_broker
    FOR EACH ROW
EXECUTE PROCEDURE update_record();

CREATE TRIGGER sync_update_record
    BEFORE UPDATE
    ON
        angel_one
    FOR EACH ROW
EXECUTE PROCEDURE update_record();

CREATE TRIGGER sync_update_record
    BEFORE UPDATE
    ON
        zerodha
    FOR EACH ROW
EXECUTE PROCEDURE update_record();

-- DROP TRIGGER sync_update_record on angel_one;
-- SELECT  event_object_table AS table_name ,trigger_name
-- FROM information_schema.triggers  
-- WHERE event_object_table ='angel_one' 
-- GROUP BY table_name , trigger_name 
-- ORDER BY table_name ,trigger_name