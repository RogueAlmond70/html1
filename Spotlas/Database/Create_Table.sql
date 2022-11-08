GRANT ALL PRIVILEGES ON DATABASE postgres TO postgres;
CREATE TABLE "MY_TABLE" (
                            id uuid,
                            name varchar,
                            website varchar,
                            coordinates geography,
                            description varchar,
                            rating float8
);