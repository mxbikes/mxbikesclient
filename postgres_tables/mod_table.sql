-- Add uuid
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Creation of mod table
CREATE TABLE IF NOT EXISTS mod (
  id uuid DEFAULT uuid_generate_v4 (),  
  name varchar(50) UNIQUE,
  description varchar(250) NOT NULL,
  id_modtypecategory uuid NOT NULL,

  releaseyear smallint NOT NULL,
  created_at timestamp  NOT NULL  DEFAULT current_timestamp,

  PRIMARY KEY (Id)
);

-- Insert data into mod table
INSERT INTO mod(name, description, id_modtypecategory, releaseyear)
VALUES ('FXR Helium - Chromatic', 'A colorful set from FXR, with the colors: gold, purple, blue.', '63964035-51ef-40a1-87e3-fa533bbbebc4', '2023'), 
('FXR Helium - Glacier', 'A colorful set from FXR, with the colors: blue, red.', '63964035-51ef-40a1-87e3-fa533bbbebc4', '2023' ), 
('FXR Helium - Ignition', 'A colorful set from FXR, with the colors: yellow, red.', '63964035-51ef-40a1-87e3-fa533bbbebc4', '2023'),
('FXR Helium - Ultra Violet', 'A colorful set from FXR, with the colors: grey, purple, blue.', '63964035-51ef-40a1-87e3-fa533bbbebc4', '2023');
