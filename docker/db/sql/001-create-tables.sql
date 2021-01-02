
CREATE TABLE users (
  user_id INT NOT NULL AUTO_INCREMENT,
  user_name varchar(50) NOT NULL,
  user_password text,
  email_address varchar(200),
  note text,
  icon_image varchar(200),
  PRIMARY KEY(user_id)
);

CREATE TABLE pictures (
  picture_id int NOT NULL AUTO_INCREMENT,
  user_id int NOT NULL,
  image_url varchar(200),
  image_note varchar(250),
  prefecture_category_cd varchar(2),
  view_category_cd varchar(2),
  like_count int,
  published_at timestamp,
  PRIMARY KEY (picture_id)
);

CREATE TABLE view_categories (
  view_category_cd varchar(3),
  view_name varchar(20),
  sort_no int,
  PRIMARY KEY (view_category_cd)
);

CREATE TABLE prefecture_categories (
  prefecture_category_cd varchar(2),
  prefecture_name varchar(4),
  sort_no int,
  PRIMARY KEY (prefecture_category_cd)
);
