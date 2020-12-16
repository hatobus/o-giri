drop table if exists ogiri;

drop table if exists user;
create table user (
  id integer primary key auto_increment,
  name varchar(32) not null
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;

drop table if exists odai;
create table odai (
  odai_id integer,
  questioner_id integer not null,
  odai varchar(128) not null,
  ogigi_id integer not null,
  next_ogiri_id varchar(36),
  published_at datetime default current_timestamp(),

  foreign key fk_odai_id (odai_id) references ogiri(odai_id),
  foreign key fk_quentioner_id (questioner_id) references user(id),
  foreign key fk_next_ogiri_id (next_ogiri_id) references ogiri(ogiri_id)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;

drop table if exists boke;
create table boke (
  id integer primary key auto_increment,
  answerer_id integer not null,
  boke varchar(128) not null,
  ogiri_id varchar(36) not null,
  published_at datetime not null,

  foreign key fk_id (id) references user(id),
  foreign key fk_answerer_id (answerer_id) references user(id),
  foreign key fk_ogiri_id (ogiri_id) references ogiri(ogiri_id)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;

drop table if exists ogiri;
create table ogiri (
  id integer primary key auto_increment,
  ogiri_id varchar(36) not null,
  odai_id integer not null,
  answer_duration integer not null,
  vote_duration integer not null,
  question_duration integer not null
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;

drop table if exists vote;
create table vote (
  id integer primary key auto_increment,
  ogiri_id integer not null,
  boke_id integer not null,
  answerer_id integer not null,

  foreign key fk_ogiri_id (ogiri_id) references ogiri(ogiri_id),
  foreign key fk_boke_id (boke_id) references boke(id),
  foreign key fk_answerer_id (answerer_id) references user(id)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;
