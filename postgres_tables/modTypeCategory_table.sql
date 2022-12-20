-- Add uuid
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Creation of modtypecategory table
CREATE TABLE IF NOT EXISTS modtypecategory (
  id uuid DEFAULT uuid_generate_v4 (),  
  name varchar(50) UNIQUE,
  id_modtype uuid NOT NULL,
  PRIMARY KEY (Id)
);

-- Insert data into modtypecategory table
INSERT INTO modtypecategory(name, id_modtype)
-- Rider uuid
VALUES ('Rider', '1393bfce-996d-4979-89d3-d10f307ae5da'), 
('Helmet', '1393bfce-996d-4979-89d3-d10f307ae5da'), 
('Jersey and Pants', '1393bfce-996d-4979-89d3-d10f307ae5da'), 
('Boots', '1393bfce-996d-4979-89d3-d10f307ae5da'), 
('Gloves', '1393bfce-996d-4979-89d3-d10f307ae5da'), 
-- Track uuid
('Enduro', '5804835c-87ed-415b-a02b-9a1b3723637d'), 
('Motorcross', '5804835c-87ed-415b-a02b-9a1b3723637d'), 
('Supercross', '5804835c-87ed-415b-a02b-9a1b3723637d'), 
('Supermoto', '5804835c-87ed-415b-a02b-9a1b3723637d'), 
-- Bike uuid
('Stock', 'bee733a6-b0f7-4291-83ce-9c7e6f31ffba'),
('Alta', 'bee733a6-b0f7-4291-83ce-9c7e6f31ffba'),
('Honda', 'bee733a6-b0f7-4291-83ce-9c7e6f31ffba'),
('Husqvarna ', 'bee733a6-b0f7-4291-83ce-9c7e6f31ffba'),
('Kawasaki ', 'bee733a6-b0f7-4291-83ce-9c7e6f31ffba'),
('KTM ', 'bee733a6-b0f7-4291-83ce-9c7e6f31ffba'),
('Suzuki ', 'bee733a6-b0f7-4291-83ce-9c7e6f31ffba'),
('TM ', 'bee733a6-b0f7-4291-83ce-9c7e6f31ffba'),
('Yamaha ', 'bee733a6-b0f7-4291-83ce-9c7e6f31ffba');
