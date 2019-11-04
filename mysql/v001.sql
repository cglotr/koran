START TRANSACTION;
CREATE TABLE quran (
  id           INT           NOT NULL AUTO_INCREMENT PRIMARY KEY,
  sura_id      INT           NOT NULL,
  verse_id     INT           NOT NULL,
  ayah         VARCHAR(2048) NOT NULL,
  UNIQUE KEY (sura_id, verse_id)
);
INSERT INTO migration (version_name) VALUES ("v001");
COMMIT;
