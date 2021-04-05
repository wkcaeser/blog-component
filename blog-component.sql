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

CREATE TABLE access_metric (
                               id bigint not null auto_increment comment 'id',
                               url varchar(128) not null comment 'url',
                               browse_cnt bigint not null comment '访问次数统计表',
                               browse_user_Cnt bigint not null comment '访问人次统计表',
                               comment varchar(128) not null comment 'url备注',
                               primary key (id),
                               unique (url)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment 'url访问统计表';