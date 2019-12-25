START TRANSACTION;
CREATE TABLE user_quran (
  user_id  INT NOT NULL,
  sura_id  INT NOT NULL,
  verse_id INT NOT NULL,
  CONSTRAINT UNIQUE KEY (user_id, sura_id, verse_id),
  CONSTRAINT FOREIGN KEY (user_id) REFERENCES user (id),
  CONSTRAINT FOREIGN KEY (sura_id, verse_id) REFERENCES quran (sura_id, verse_id)
);
INSERT INTO migration (version_name) VALUES ("v004");
COMMIT;
