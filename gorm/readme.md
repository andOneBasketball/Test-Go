CREATE DATABASE IF NOT EXISTS gorm_research;
drop database gorm_research;
CREATE USER 'gorm_research'@'localhost' IDENTIFIED BY 'gorm_research1234567890';
GRANT ALL PRIVILEGES ON gorm_research.* TO 'gorm_research'@'localhost';

INSERT INTO user_infos(name, gender, hobby) VALUES("ff", "girl", "tennis");