
CREATE OR REPLACE FUNCTION create_db(dbname text)
  RETURNS void AS
$func$
BEGIN
IF EXISTS (SELECT 1 FROM pg_database WHERE datname = dbname) THEN
   RAISE NOTICE 'Database already exists';
ELSE
   PERFORM dblink_exec('dbname=' || current_database()
                     , 'CREATE DATABASE ' || quote_ident(dbname));
END IF;

END
$func$ LANGUAGE plpgsql;



CREATE OR REPLACE FUNCTION drop_db(dbname text)
  RETURNS void AS
$func$
BEGIN
IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = dbname) THEN
   RAISE NOTICE 'There is no database with such name';
ELSE
   PERFORM dblink_exec('dbname=' || current_database()
                     , 'DROP DATABASE ' || quote_ident(dbname));
END IF;

END
$func$ LANGUAGE plpgsql;


