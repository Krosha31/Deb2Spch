DROP FUNCTION IF EXISTS create_all_tables();

CREATE FUNCTION create_all_tables() RETURNS void AS $$
BEGIN
    CREATE TABLE IF NOT EXISTS users (
        id serial PRIMARY KEY,
        "login" varchar(200),  -- Изменено с varchar[200] на varchar(200)
        password_hash varchar(200),
        subscribtion_id integer DEFAULT 0,
        registration_date date DEFAULT CURRENT_DATE
    );

    CREATE TABLE IF NOT EXISTS subscriptions (  -- Исправлено "subscribtions" на "subscriptions"
        id serial PRIMARY KEY, 
        date_start date DEFAULT CURRENT_DATE,
        duration_in_months integer  -- Удалена запятая в конце
    );

    CREATE TABLE IF NOT EXISTS logs (
        log_id serial PRIMARY KEY,
        "action" text,
        "date" date DEFAULT CURRENT_DATE,
        user_id integer,
        subscribtion_id integer
    );

    ALTER TABLE users ADD FOREIGN KEY (subscribtion_id) REFERENCES subscriptions(id);  -- Исправлено "subscribtions" на "subscriptions"

    ALTER TABLE logs ADD FOREIGN KEY (user_id) REFERENCES users(id);  -- Исправлено "partner user" на "users"

    ALTER TABLE logs ADD FOREIGN KEY (subscribtion_id) REFERENCES subscriptions(id);  -- Исправлено "subscribtions" на "subscriptions"

END;  -- Добавлен блок BEGIN...END для функции
$$ LANGUAGE plpgsql;  -- Изменен язык на plpgsql, так как функция содержит несколько операторов
