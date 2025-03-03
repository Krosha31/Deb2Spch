CREATE OR REPLACE FUNCTION add_user(input_email varchar(200), input_password_hash varchar(200))
RETURNS void AS $func$
BEGIN
    INSERT INTO users (email, password_hash) VALUES (input_email, input_password_hash);
END;
$func$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION add_suncription(duration integer)
RETURNS void AS $func$
BEGIN
    INSERT INTO subscriptions (duration_in_months) VALUES (duration);
END;
$func$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION get_user_by_login(input_email varchar(200))
RETURNS TABLE(email varchar(200),
              password_hash varchar(200),
              subscribtion_id integer, 
              registration_date date) AS $func$
BEGIN
    RETURN QUERY
    SELECT users.email, users.password_hash, users.subscribtion_id, users.registration_date
    FROM users 
    WHERE users.email = input_email;
END;
$func$ LANGUAGE plpgsql;
