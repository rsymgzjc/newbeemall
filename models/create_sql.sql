drop table if exists user;

create table user(
         id bigint(20) not null auto_increment,
         user_id bigint(20) not null ,
         username varchar(64) collate utf8mb4_general_ci not null,
         password varchar(64) collate utf8mb4_general_ci not null,
         introduction varchar(128) collate utf8mb4_general_ci not null ,
         email varchar(64) null collate utf8mb4_general_ci,
         gender tinyint(4) not null default '0',
         create_time timestamp null default current_timestamp,
         update_time timestamp null default current_timestamp on update current_timestamp,
         PRIMARY KEY  (id),
         UNIQUE KEY idx_username (username) USING BTREE ,
         UNIQUE KEY idx_user_id (user_id) USING BTREE
)engine =InnoDB DEFAULT CHARSET =utf8mb4 COLLATE=utf8mb4_general_ci;

alter table user modify introduction varchar(128) collate utf8mb4_general_ci not null default '';

drop table if exists user_address;

create table user_address(
         address_id bigint(20) not null auto_increment,
         user_id bigint(20) not null ,
         username varchar(64) collate utf8mb4_general_ci not null,
         userphone varchar(64) collate utf8mb4_general_ci not null,
         defaultflag tinyint(4) not null default '0',
         provincename varchar(64) collate utf8mb4_general_ci not null,
         cityname varchar(64) collate utf8mb4_general_ci not null,
         regionname varchar(64) collate utf8mb4_general_ci not null,
         detailaddress varchar(64) collate utf8mb4_general_ci not null,
         create_time timestamp null default current_timestamp,
         update_time timestamp null default current_timestamp on update current_timestamp,
         PRIMARY KEY  (address_id),
         UNIQUE KEY idx_username (username) USING BTREE ,
         UNIQUE KEY idx_user_id (user_id) USING BTREE
)engine =InnoDB DEFAULT CHARSET =utf8mb4 COLLATE=utf8mb4_general_ci;

alter table user_address drop index idx_user_id;
drop table if exists admin_user;
create table admin_user(
         id bigint(20) not null auto_increment,
         admin_id bigint(20) not null ,
         adminname varchar(64) collate utf8mb4_general_ci not null,
         password varchar(64) collate utf8mb4_general_ci not null,
         create_time timestamp null default current_timestamp,
         update_time timestamp null default current_timestamp on update current_timestamp,
         PRIMARY KEY  (id),
         UNIQUE KEY idx_username (adminname) USING BTREE ,
         UNIQUE KEY idx_user_id (admin_id) USING BTREE
)engine =InnoDB DEFAULT CHARSET =utf8mb4 COLLATE=utf8mb4_general_ci;

alter table user add lockflag tinyint(4) not null default 0;

create table goods_info(
       goods_id bigint(20) not null auto_increment,
       goodsname varchar(64) collate utf8mb4_general_ci not null,
       goodsintro varchar(128) collate utf8mb4_general_ci not null,
       goodscategory_id bigint(20) not null ,
       goodscoverimg varchar(128) null ,
       goodscarousel varchar(256) null ,
       goodsdetail varchar(128) not null ,
       originprice int(10) not null ,
       sellingprice int(10) not null ,
       stocknum int(10) not null ,
       tag varchar(20) not null ,
       goodssellstatus tinyint(4) not null  default '0',
       create_time timestamp null default current_timestamp,
       update_time timestamp null default current_timestamp on update current_timestamp,
       PRIMARY KEY  (goods_id),
       UNIQUE KEY idx_username (goodsname) USING BTREE ,
       UNIQUE KEY idx_user_id (goodscategory_id) USING BTREE
)engine =InnoDB DEFAULT CHARSET =utf8mb4 COLLATE=utf8mb4_general_ci;

alter table goods_info drop index idx_user_id;
alter table goods_info drop index idx_username;

create table goods_category(
       category_id bigint(20) not null auto_increment,
       categoryname varchar(64) collate utf8mb4_general_ci not null,
       categorylevel int(10) not null,
       parentid int(10) not null,
       categoryrank int(10) null,
       isdeleted tinyint(4) not null default '0',
       create_time timestamp null default current_timestamp,
       update_time timestamp null default current_timestamp on update current_timestamp,
       PRIMARY KEY  (category_id)
)engine =InnoDB DEFAULT CHARSET =utf8mb4 COLLATE=utf8mb4_general_ci;
alter table goods_category modify categoryrank int(10) not null default '0';

create table carousel(
       carousel_id bigint(20) not null auto_increment,
       carouselurl varchar(64) collate utf8mb4_general_ci not null,
       redirecturl varchar(64) collate utf8mb4_general_ci not null,
       carouselrank int(10) null,
       isdeleted tinyint(4) not null default '0',
       create_time timestamp null default current_timestamp,
       update_time timestamp null default current_timestamp on update current_timestamp,
       PRIMARY KEY  (carousel_id)
)engine =InnoDB DEFAULT CHARSET =utf8mb4 COLLATE=utf8mb4_general_ci;

create table indexconfig(
         config_id bigint(20) not null auto_increment,
         configname varchar(40) not null ,
         configtype tinyint(4) collate utf8mb4_general_ci not null,
         goods_id bigint(20) not null default '0',
         redirecturl varchar(64) collate utf8mb4_general_ci not null,
         configrank int(10) null,
         isdeleted tinyint(4) not null default '0',
         create_time timestamp null default current_timestamp,
         update_time timestamp null default current_timestamp on update current_timestamp,
         PRIMARY KEY  (config_id)
)engine =InnoDB DEFAULT CHARSET =utf8mb4 COLLATE=utf8mb4_general_ci;

create table mallorder(
        order_id bigint(20) not null ,
        ordernum varchar(20) not null ,
        user_id bigint(20) not null ,
        totalprice int(10) not null ,
        paystatus tinyint(4) not null ,
        paytype tinyint(4) not null ,
        pay_time timestamp null,
        configrank int(10) null,
        orderstatus tinyint(4) not null ,
        extrainfo varchar(100) not null  default '',
        isdeleted tinyint(4) not null default '0',
        create_time timestamp null default current_timestamp,
        update_time timestamp null default current_timestamp on update current_timestamp,
        PRIMARY KEY  (order_id)
)engine =InnoDB DEFAULT CHARSET =utf8mb4 COLLATE=utf8mb4_general_ci;

alter table mallorder modify order_id bigint(20) not null auto_increment ;
alter table mallorder modify ordernum bigint(20) not null;
alter table mallorder modify order_id bigint(20) not null;

create table shoppingcart(
      cart_id bigint(20) not null auto_increment,
      user_id bigint(20) not null ,
      goods_id bigint(20) not null,
      goods_count int(10) not null ,
      isdeleted tinyint(4) not null default '0',
      create_time timestamp null default current_timestamp,
      update_time timestamp null default current_timestamp on update current_timestamp,
      PRIMARY KEY  (cart_id)
)engine =InnoDB DEFAULT CHARSET =utf8mb4 COLLATE=utf8mb4_general_ci;

create table orderitem(
        orderitem_id bigint(20) not null auto_increment,
        order_id bigint(20) not null ,
        goods_id bigint(20) not null,
        goodsname varchar(64) collate utf8mb4_general_ci not null,
        goodscoverimg varchar(128) null ,
        sellingprice int(10) not null ,
        goods_count int(10) not null ,
        create_time timestamp null default current_timestamp,
        PRIMARY KEY  (orderitem_id)
)engine =InnoDB DEFAULT CHARSET =utf8mb4 COLLATE=utf8mb4_general_ci;
