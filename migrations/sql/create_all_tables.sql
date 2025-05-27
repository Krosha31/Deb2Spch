DROP FUNCTION IF EXISTS create_all_tables();

CREATE FUNCTION create_all_tables() RETURNS void AS $$
BEGIN
    CREATE TABLE IF NOT EXISTS users (
        email varchar(200) PRIMARY KEY,
        password_hash varchar(200),
        subscribtion_id integer DEFAULT 1,
        registration_date date DEFAULT CURRENT_DATE,
        free_minutes integer DEFAULT 10
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
        user_id varchar(200) NOT NULL,
        subscribtion_id integer
    );

    CREATE TABLE IF NOT EXISTS request (
        id SERIAL PRIMARY KEY,
        user_id varchar(200) NOT NULL,
        time TIMESTAMP NOT NULL DEFAULT now(),
        input TEXT NOT NULL
    );

    CREATE TABLE IF NOT EXISTS request_results (
        id SERIAL PRIMARY KEY,
        request_id INT NOT NULL,
        "output" TEXT NOT NULL
    );

    ALTER TABLE users ADD FOREIGN KEY (subscribtion_id) REFERENCES subscriptions(id);  
    ALTER TABLE logs ADD FOREIGN KEY (user_id) REFERENCES users(email);  

    ALTER TABLE logs ADD FOREIGN KEY (subscribtion_id) REFERENCES subscriptions(id);  

    ALTER TABLE request ADD FOREIGN KEY (user_id) REFERENCES users(email);
    ALTER TABLE request_results ADD FOREIGN KEY (request_id) REFERENCES request(id) ON DELETE CASCADE;

END;  
$$ LANGUAGE plpgsql; 
