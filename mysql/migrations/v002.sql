START TRANSACTION;
CREATE TABLE translation (
  id           INT           NOT NULL AUTO_INCREMENT PRIMARY KEY,
  sura_id      INT           NOT NULL,
  verse_id     INT           NOT NULL,
  ayah         VARCHAR(2048) NOT NULL,
  name         VARCHAR(16)   NOT NULL,
  UNIQUE KEY (sura_id, verse_id, name)
);
INSERT INTO migration (version_name) VALUES ("v002");
COMMIT;
