CREATE TABLE talents (
    id bigserial PRIMARY KEY,
    full_name varchar NOT NULL,
    gender varchar,
    year_of_birth integer ,
    phone varchar,
    email varchar,
    applied_position varchar,
    level varchar,
    department varchar,
    project varchar,
    cv varchar,
    criteria varchar,
    scheduled_interview timestamp with time zone,
    interview_result varchar,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);

CREATE TABLE users (
    id bigserial PRIMARY KEY,
    username varchar NOT NULL,
    password varchar NOT NULL,
    Name varchar
);

CREATE TABLE levels (
    id bigserial PRIMARY KEY,
    code varchar NOT NULL,
    name varchar NOT NULL
);

INSERT INTO levels (code, name) VALUES ('Junior', 'Junior');
INSERT INTO levels (code, name) VALUES ('Middle', 'Middle');
INSERT INTO levels (code, name) VALUES ('Senior', 'Senior');
INSERT INTO levels (code, name) VALUES ('TechLead', 'Tech Lead');