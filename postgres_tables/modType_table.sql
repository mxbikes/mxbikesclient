-- Add uuid
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Creation of modtype table
CREATE TABLE IF NOT EXISTS modtype (
  id uuid DEFAULT uuid_generate_v4 (),  
  name varchar(50) UNIQUE,
  PRIMARY KEY (Id)
);

-- Insert data into modtype table
INSERT INTO modtype(name)
VALUES ('Rider'), ('Bike'), ('Track');
