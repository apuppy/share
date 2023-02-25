CREATE DATABASE `blog`
/*!40100 DEFAULT CHARACTER SET utf8mb4 */
;
------------------------------ bookmark ------------------------------
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
------------------------------ proxy server ------------------------------
-- proxy
CREATE TABLE proxy (
    id int primary key auto_increment,
    name varchar(32) not null comment 'name',
    type varchar(16) not null comment 'type: vmess, ss, ssr, trojan',
    server varchar(128) not null comment 'server address',
    port int unsigned not null comment 'proxy port',
    uuid varchar(64) not null default '' comment 'uuid',
    alter_id int unsigned not null default 0 comment 'alter_id',
    cipher varchar(64) unsigned not null default '' comment 'cipher',
    password varchar(64) unsigned not null default '' comment 'password',
    udp tinyint(1) unsigned not null default 0 comment 'udp, false-0, true-1',
    created_by varchar(100) NOT NULL DEFAULT '' COMMENT 'creator',
    updated_by varchar(100) NOT NULL DEFAULT '' COMMENT 'updater',
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create datetime',
    updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'update datetime',
    deleted tinyint(1) DEFAULT 0 COMMENT 'delete status: normal-0, deleted-1'
) engine = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8_unicode_ci comment 'proxy';
-- proxy_group
CREATE TABLE proxy_group (
    id int primary key auto_increment,
    name varchar(32) not null comment 'name',
    created_by varchar(100) NOT NULL DEFAULT '' COMMENT 'creator',
    updated_by varchar(100) NOT NULL DEFAULT '' COMMENT 'updater',
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create datetime',
    updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'update datetime',
    deleted tinyint(1) DEFAULT 0 COMMENT 'delete status: normal-0, deleted-1'
) engine = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8_unicode_ci comment 'proxy group';
-- proxy_group_proxy_mapping
CREATE TABLE proxy_group_proxy_mapping (
    id int primary key auto_increment,
    group_id int unsigned not null comment 'proxy_group id',
    proxy_id int unsigned not null comment 'proxy id',
    created_by varchar(100) NOT NULL DEFAULT '' COMMENT 'creator',
    updated_by varchar(100) NOT NULL DEFAULT '' COMMENT 'updater',
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create datetime',
    updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'update datetime',
    deleted tinyint(1) DEFAULT 0 COMMENT 'delete status: normal-0, deleted-1'
) engine = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8_unicode_ci comment 'proxy and proxy_group mapping';
-- proxy_rule
-- match_type: DOMAIN, DOMAIN-KEYWORD, DOMAIN-SUFFIX, GEOIP, IP-CIDR, IP-CIDR6, MATCH
-- proxy_type: PROXY, DERECT, REJECT, MISSED
CREATE TABLE proxy_rule (
    id int primary key auto_increment,
    match_type char(16) not null comment 'match_type: DOMAIN, DOMAIN-KEYWORD, DOMAIN-SUFFIX, GEOIP, IP-CIDR, IP-CIDR6, MATCH',
    proxy_type char(16) not null comment 'proxy_type: PROXY, DERECT, REJECT, MISSED',
    dst_addr varchar(128) not null comment 'destnation address',
    updated_by varchar(100) NOT NULL DEFAULT '' COMMENT 'updater',
    created_by varchar(100) NOT NULL DEFAULT '' COMMENT 'creator',
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create datetime',
    updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'update datetime',
    deleted tinyint(1) DEFAULT 0 COMMENT 'delete status: normal-0, deleted-1'
) engine = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8_unicode_ci comment 'proxy rule';