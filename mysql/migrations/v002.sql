START TRANSACTION;
CREATE TABLE translation (
  id           INT           NOT NULL AUTO_INCREMENT PRIMARY KEY,
  sura_id      INT           NOT NULL,
  verse_id     INT           NOT NULL,
  name         VARCHAR(16)   NOT NULL,
  translation  VARCHAR(2048) NOT NULL,
  CONSTRAINT UNIQUE KEY (sura_id, verse_id, name),
  CONSTRAINT FOREIGN KEY (sura_id, verse_id) REFERENCES quran (sura_id, verse_id)
);
INSERT INTO migration (version_name) VALUES ("v002");
COMMIT;
