ALTER DATABASE challenge_appdb OWNER TO spuser;

CREATE SCHEMA challenge AUTHORIZATION spuser;
GRANT CONNECT ON DATABASE challenge_appdb TO spuser;
GRANT ALL ON DATABASE challenge_appdb TO spuser WITH GRANT OPTION;
GRANT USAGE ON SCHEMA challenge TO spuser WITH GRANT OPTION;
GRANT ALL ON SCHEMA challenge TO spuser WITH GRANT OPTION;

CREATE SEQUENCE account_id_seq;
CREATE TABLE challenge.account (
       id bigint NOT NULL DEFAULT NEXTVAL('account_id_seq'),
       name character varying(100) NOT NULL,
       cpf character varying(11) NOT NULL UNIQUE,
       secret character varying(256) NOT NULL,
       balance numeric NOT NULL,
       created_at timestamp with time zone NOT NULL,
       PRIMARY KEY (id)
    );

INSERT INTO challenge.account (name, cpf, secret, balance, created_at) VALUES ('Joe Doe', '11122233344', '$2a$08$0tzNbzLdjRm91KlHzLdmzOIE8rHNsk3bv1VLXn6SBY6ONgxJkRefq', 200, '2020-10-19 10:23:54+02');
INSERT INTO challenge.account (name, cpf, secret, balance, created_at) VALUES ('Sarah Connor', '12345678900', '$2a$08$7atD1ViDYZtbrQzTzptc4.cvxvGr.9ZR6YRmikwcTqT2Dask9N.JK', 1000, '2021-10-19 10:23:54+02');

CREATE SEQUENCE transfer_id_seq;
CREATE TABLE challenge.transfer (
    id bigint NOT NULL DEFAULT NEXTVAL('transfer_id_seq'),
    account_origin_id bigint NOT NULL,
    account_destination_id bigint NOT NULL,
    amount numeric NOT NULL,
    created_at timestamp with time zone,
    PRIMARY KEY (id),
    CONSTRAINT account_destination_id_fkey FOREIGN KEY (account_destination_id) REFERENCES challenge.account(id),
    CONSTRAINT account_origin_id_fkey FOREIGN KEY (account_origin_id) REFERENCES challenge.account(id)
);