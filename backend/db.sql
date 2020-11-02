CREATE DATABASE `blog` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

create table link_category(
    id int unsigned primary key auto_increment,
    name varchar(30) not null default '',
    sort tinyint(3) unsigned not null default 0 comment 'sort the smaller the number, the higher priority',
    is_delete tinyint(1) unsigned not null default 0 comment '0:not delete，1:deleted',
    created_at datetime not null default current_timestamp,
    updated_at datetime,
    deleted_at datetime
) engine = InnoDB default charset = utf8mb4 comment 'link category';

create table link(
    id int unsigned primary key auto_increment,
    title varchar(100) not null default '' comment 'title',
    url varchar(300) not null default '' comment 'URL',
    is_delete tinyint(1) unsigned not null default 0 comment '0:not delete，1:deleted',
    created_at datetime not null default current_timestamp,
    updated_at datetime,
    deleted_at datetime
) engine = InnoDB default charset = utf8mb4 comment 'collected links';