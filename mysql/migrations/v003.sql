START TRANSACTION;
CREATE TABLE user (
  id    INT          NOT NULL AUTO_INCREMENT PRIMARY KEY,
  uid   VARCHAR(100) NOT NULL,
  token VARCHAR(100) NOT NULL,
  CONSTRAINT UNIQUE KEY (uid)
);
INSERT INTO migration (version_name) VALUES ("v003");
COMMIT;
