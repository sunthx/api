create database if not exists core;

use core;

create table if not exists files (
  id int auto_increment,
  file_id char(50) not null,
  file_name char(200),
  file_type char(50),
  file_size double default 0,
  file_create_date date,
  primary key (id)
)