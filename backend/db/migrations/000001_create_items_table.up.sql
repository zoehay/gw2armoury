CREATE TABLE items (
  id         INT NOT NULL,
  chat_link VARCHAR(128), 
  name      VARCHAR(128),
  icon     VARCHAR(255),
  description      VARCHAR(255),
  type VARCHAR(128) NOT NULL,
  rarity VARCHAR(128),
  level INT,
  vendor_value INT,
  default_skin INT,
  flags VARCHAR(255),
  game_types VARCHAR(255),
  restrictions VARCHAR(255),
  upgrades_into VARCHAR(255),
  upgrades_from VARCHAR(255),
  details VARCHAR(255),
  PRIMARY KEY (`id`)
);