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

-- CREATE OR REPLACE FUNCTION insert_request(
--     p_user_id varchar,
--     p_input TEXT
-- ) RETURNS INTEGER AS $$
-- DECLARE
--     new_id INTEGER;
-- BEGIN
--     INSERT INTO request (user_id, input)
--     VALUES (p_user_id, p_input)
--     RETURNING id INTO new_id;
    
--     RETURN new_id;
-- END;
-- $$ LANGUAGE plpgsql;

-- CREATE OR REPLACE FUNCTION get_requests_by_user(input_user_id varchar)
-- RETURNS TABLE(
--     id integer,
--     user_id varchar,
--     time timestamp,
--     input text
-- ) AS $func$
-- BEGIN
--     RETURN QUERY
--     SELECT id, user_id, time, input
--     FROM request
--     WHERE user_id = input_user_id
--     ORDER BY time DESC;
-- END;
-- $func$ LANGUAGE plpgsql;