drop table if exists supported_broker;
drop table if exists angel_one;
drop table if exists zerodha;
drop table if exists idirect;
drop table if exists mf_central;
drop table if exists creds;

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

create table supported_sources(
    id serial primary key,
    source_name text unique not null,
    source_type text,
    holdings_sync bool not null default FALSE,
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

create table mf_central(
    id serial primary key,
    folio text not null,
    scheme_name text not null,
    isin text not null,
    quantity float not null,
    price float not null,
    cost_price float,
    curr_price float,
    created_on timestamp default now(),
    updated_on timestamp default now()
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
        supported_sources
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

CREATE TRIGGER sync_update_record
    BEFORE UPDATE
    ON
        mf_central
    FOR EACH ROW
EXECUTE PROCEDURE update_record();

-- DROP TRIGGER sync_update_record on angel_one;
-- SELECT  event_object_table AS table_name ,trigger_name
-- FROM information_schema.triggers  
-- WHERE event_object_table ='angel_one' 
-- GROUP BY table_name , trigger_name 
-- ORDER BY table_name ,trigger_name

-- seed data
insert into supported_sources(source_name,source_type,holdings_sync) values
('zerodha','broker','f'),
('angelone','broker','f'),
('idirect','broker','f'),
('mfcentral','mutualfund','f');