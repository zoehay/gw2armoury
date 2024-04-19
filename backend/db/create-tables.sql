DROP TABLE IF EXISTS item;

CREATE TABLE item (
  id         INT NOT NULL,
  chat_link VARCHAR(128), 
  name      VARCHAR(128),
  icon     VARCHAR(255),
  description      VARCHAR(255),
  type VARCHAR(128),
  rarity VARCHAR(128),
  level INT,
  vendor_value
  PRIMARY KEY (`id`)
);

INSERT INTO item   
    (id, name, icon, description)
VALUES
    {28445, "Strong Soft Wood Longbow of Fire", "https://render.guildwars2.com/file/C6110F52DF5AFE0F00A56F9E143E9732176DDDE9/65015.png", ""},
    {12452, "Omnomberry Bar", "https://render.guildwars2.com/file/6BD5B65FBC6ED450219EC86DD570E59F4DA3791F/433643.png", ""},