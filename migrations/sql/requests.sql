CREATE OR REPLACE FUNCTION add_user(input_login varchar(200), input_password_hash varchar(200))
RETURNS void AS $$
    INSERT INTO users VALUES (input_login, input_password_hash);
$$ LANGUAGE sql;

CREATE OR REPLACE FUNCTION get_user_by_login(input_login varchar(200))
RETURNS TABLE(id integer,
        "login" varchar(200),
        password_hash varchar(200),
        subscribtion_id integer,
    ) AS $$
    SELECT id, "login" password_hash, subscribtion_id FROM users WHERE users.login == input_login;
$$ LANGUAGE sql;