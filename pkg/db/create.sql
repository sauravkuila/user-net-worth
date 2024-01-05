drop table if exists angel_one;
drop table if exists supported_broker;

create table angel_one(
    id serial primary key,
    symbol text not null,
    isin text not null,
    quantity int not null,
    price float not null,
    created_on date default now(),
    updated_on date default now()
);

create table supported_broker(
    id serial primary key,
    broker_name text not null,
    created_on date default now(),
    updated_on date default now()
);

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

-- DROP TRIGGER sync_update_record on angel_one;
-- SELECT  event_object_table AS table_name ,trigger_name
-- FROM information_schema.triggers  
-- WHERE event_object_table ='angel_one' 
-- GROUP BY table_name , trigger_name 
-- ORDER BY table_name ,trigger_name