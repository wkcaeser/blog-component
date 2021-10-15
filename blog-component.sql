create database blog_component;

use blog_component;

CREATE Table user (
                      id bigint not null auto_increment comment '用户id',
                      create_time datetime not null comment '创建时间',
                      modify_time datetime not null comment '更新时间',
                      username varchar(32) not null comment '用户名称',
                      nick varchar(32) not null comment '昵称',
                      email varchar(64) not null comment '邮箱',
                      password varchar(64) not null comment '密码',
                      icon longtext not null comment '头像',
                      status int not null comment '状态 1：正常',
                      primary key (id),
                      unique (username),
                      unique (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment '用户表';

CREATE TABLE user_token (
                            id bigint not null auto_increment comment 'id',
                            token varchar(128) not null comment 'token val',
                            user_id bigint not null comment '用户id',
                            create_time datetime not null comment '创建时间',
                            expire_time datetime not null comment '过期时间',
                            primary key (id),
                            unique (token)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment '用户token表';

CREATE TABLE access_log (
                               id bigint not null auto_increment comment 'id',
                               user_uuid varchar(32) not null comment 'uuid',
                               uri varchar(128) not null comment 'uri',
                               site varchar(128) not null comment 'site',
                               req_uri varchar(256) comment 'req_uri',
                               ip varchar(128) comment 'source_ip',
                               request_log longtext not null comment '请求信息',
                               primary key (id),
                               index idx_uuid (user_uuid),
                               index idx_uri(uri)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment 'url访问统计表';


SELECT count(distinct user_uuid) as PeopleTotal from access_log;