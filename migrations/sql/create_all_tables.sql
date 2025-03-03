DROP FUNCTION IF EXISTS create_all_tables();

CREATE FUNCTION create_all_tables() RETURNS void AS $$
BEGIN
    CREATE TABLE IF NOT EXISTS users (
        email varchar(200) PRIMARY KEY,
        password_hash varchar(200),
        subscribtion_id integer DEFAULT 1,
        registration_date date DEFAULT CURRENT_DATE
    );

    CREATE TABLE IF NOT EXISTS subscriptions (  
        id serial PRIMARY KEY, 
        date_start date DEFAULT CURRENT_DATE,
        duration_in_months integer  
    );

    CREATE TABLE IF NOT EXISTS logs (
        log_id serial PRIMARY KEY,
        "action" text,
        "date" date DEFAULT CURRENT_DATE,
        user_id varchar(200),
        subscribtion_id integer
    );

    ALTER TABLE users ADD FOREIGN KEY (subscribtion_id) REFERENCES subscriptions(id);  
    ALTER TABLE logs ADD FOREIGN KEY (user_id) REFERENCES users(email);  

    ALTER TABLE logs ADD FOREIGN KEY (subscribtion_id) REFERENCES subscriptions(id);  

END;  
$$ LANGUAGE plpgsql; 
