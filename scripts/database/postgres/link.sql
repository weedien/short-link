CREATE TABLE "group_0"
(
    "id"          int8 NOT NULL,
    "gid"         varchar(32),
    "name"        varchar(64),
    "username"    varchar(256),
    "sort_order"  int4,
    "create_time" timestamp,
    "update_time" timestamp,
    "del_flag"    int2,
    PRIMARY KEY ("id")
);
CREATE INDEX "idx_username" ON "group_0" USING btree (
                                                        "username" ASC
    );
COMMENT ON COLUMN "group_0"."id" IS 'ID';
COMMENT ON COLUMN "group_0"."gid" IS '分组标识';
COMMENT ON COLUMN "group_0"."name" IS '分组名称';
COMMENT ON COLUMN "group_0"."username" IS '创建分组用户名';
COMMENT ON COLUMN "group_0"."sort_order" IS '分组排序';
COMMENT ON COLUMN "group_0"."create_time" IS '创建时间';
COMMENT ON COLUMN "group_0"."update_time" IS '修改时间';
COMMENT ON COLUMN "group_0"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "group_1"
(
    "id"          int8 NOT NULL,
    "gid"         varchar(32),
    "name"        varchar(64),
    "username"    varchar(256),
    "sort_order"  int4,
    "create_time" timestamp,
    "update_time" timestamp,
    "del_flag"    int2,
    CONSTRAINT "_copy_72" PRIMARY KEY ("id")
);
CREATE INDEX "idx_username_copy_15" ON "group_1" USING btree (
                                                                "username" ASC
    );
COMMENT ON COLUMN "group_1"."id" IS 'ID';
COMMENT ON COLUMN "group_1"."gid" IS '分组标识';
COMMENT ON COLUMN "group_1"."name" IS '分组名称';
COMMENT ON COLUMN "group_1"."username" IS '创建分组用户名';
COMMENT ON COLUMN "group_1"."sort_order" IS '分组排序';
COMMENT ON COLUMN "group_1"."create_time" IS '创建时间';
COMMENT ON COLUMN "group_1"."update_time" IS '修改时间';
COMMENT ON COLUMN "group_1"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "group_10"
(
    "id"          int8 NOT NULL,
    "gid"         varchar(32),
    "name"        varchar(64),
    "username"    varchar(256),
    "sort_order"  int4,
    "create_time" timestamp,
    "update_time" timestamp,
    "del_flag"    int2,
    CONSTRAINT "_copy_71" PRIMARY KEY ("id")
);
CREATE INDEX "idx_username_copy_14" ON "group_10" USING btree (
                                                                 "username" ASC
    );
COMMENT ON COLUMN "group_10"."id" IS 'ID';
COMMENT ON COLUMN "group_10"."gid" IS '分组标识';
COMMENT ON COLUMN "group_10"."name" IS '分组名称';
COMMENT ON COLUMN "group_10"."username" IS '创建分组用户名';
COMMENT ON COLUMN "group_10"."sort_order" IS '分组排序';
COMMENT ON COLUMN "group_10"."create_time" IS '创建时间';
COMMENT ON COLUMN "group_10"."update_time" IS '修改时间';
COMMENT ON COLUMN "group_10"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "group_11"
(
    "id"          int8 NOT NULL,
    "gid"         varchar(32),
    "name"        varchar(64),
    "username"    varchar(256),
    "sort_order"  int4,
    "create_time" timestamp,
    "update_time" timestamp,
    "del_flag"    int2,
    CONSTRAINT "_copy_70" PRIMARY KEY ("id")
);
CREATE INDEX "idx_username_copy_13" ON "group_11" USING btree (
                                                                 "username" ASC
    );
COMMENT ON COLUMN "group_11"."id" IS 'ID';
COMMENT ON COLUMN "group_11"."gid" IS '分组标识';
COMMENT ON COLUMN "group_11"."name" IS '分组名称';
COMMENT ON COLUMN "group_11"."username" IS '创建分组用户名';
COMMENT ON COLUMN "group_11"."sort_order" IS '分组排序';
COMMENT ON COLUMN "group_11"."create_time" IS '创建时间';
COMMENT ON COLUMN "group_11"."update_time" IS '修改时间';
COMMENT ON COLUMN "group_11"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "group_12"
(
    "id"          int8 NOT NULL,
    "gid"         varchar(32),
    "name"        varchar(64),
    "username"    varchar(256),
    "sort_order"  int4,
    "create_time" timestamp,
    "update_time" timestamp,
    "del_flag"    int2,
    CONSTRAINT "_copy_69" PRIMARY KEY ("id")
);
CREATE INDEX "idx_username_copy_12" ON "group_12" USING btree (
                                                                 "username" ASC
    );
COMMENT ON COLUMN "group_12"."id" IS 'ID';
COMMENT ON COLUMN "group_12"."gid" IS '分组标识';
COMMENT ON COLUMN "group_12"."name" IS '分组名称';
COMMENT ON COLUMN "group_12"."username" IS '创建分组用户名';
COMMENT ON COLUMN "group_12"."sort_order" IS '分组排序';
COMMENT ON COLUMN "group_12"."create_time" IS '创建时间';
COMMENT ON COLUMN "group_12"."update_time" IS '修改时间';
COMMENT ON COLUMN "group_12"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "group_13"
(
    "id"          int8 NOT NULL,
    "gid"         varchar(32),
    "name"        varchar(64),
    "username"    varchar(256),
    "sort_order"  int4,
    "create_time" timestamp,
    "update_time" timestamp,
    "del_flag"    int2,
    CONSTRAINT "_copy_68" PRIMARY KEY ("id")
);
CREATE INDEX "idx_username_copy_11" ON "group_13" USING btree (
                                                                 "username" ASC
    );
COMMENT ON COLUMN "group_13"."id" IS 'ID';
COMMENT ON COLUMN "group_13"."gid" IS '分组标识';
COMMENT ON COLUMN "group_13"."name" IS '分组名称';
COMMENT ON COLUMN "group_13"."username" IS '创建分组用户名';
COMMENT ON COLUMN "group_13"."sort_order" IS '分组排序';
COMMENT ON COLUMN "group_13"."create_time" IS '创建时间';
COMMENT ON COLUMN "group_13"."update_time" IS '修改时间';
COMMENT ON COLUMN "group_13"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "group_14"
(
    "id"          int8 NOT NULL,
    "gid"         varchar(32),
    "name"        varchar(64),
    "username"    varchar(256),
    "sort_order"  int4,
    "create_time" timestamp,
    "update_time" timestamp,
    "del_flag"    int2,
    CONSTRAINT "_copy_67" PRIMARY KEY ("id")
);
CREATE INDEX "idx_username_copy_10" ON "group_14" USING btree (
                                                                 "username" ASC
    );
COMMENT ON COLUMN "group_14"."id" IS 'ID';
COMMENT ON COLUMN "group_14"."gid" IS '分组标识';
COMMENT ON COLUMN "group_14"."name" IS '分组名称';
COMMENT ON COLUMN "group_14"."username" IS '创建分组用户名';
COMMENT ON COLUMN "group_14"."sort_order" IS '分组排序';
COMMENT ON COLUMN "group_14"."create_time" IS '创建时间';
COMMENT ON COLUMN "group_14"."update_time" IS '修改时间';
COMMENT ON COLUMN "group_14"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "group_15"
(
    "id"          int8 NOT NULL,
    "gid"         varchar(32),
    "name"        varchar(64),
    "username"    varchar(256),
    "sort_order"  int4,
    "create_time" timestamp,
    "update_time" timestamp,
    "del_flag"    int2,
    CONSTRAINT "_copy_66" PRIMARY KEY ("id")
);
CREATE INDEX "idx_username_copy_9" ON "group_15" USING btree (
                                                                "username" ASC
    );
COMMENT ON COLUMN "group_15"."id" IS 'ID';
COMMENT ON COLUMN "group_15"."gid" IS '分组标识';
COMMENT ON COLUMN "group_15"."name" IS '分组名称';
COMMENT ON COLUMN "group_15"."username" IS '创建分组用户名';
COMMENT ON COLUMN "group_15"."sort_order" IS '分组排序';
COMMENT ON COLUMN "group_15"."create_time" IS '创建时间';
COMMENT ON COLUMN "group_15"."update_time" IS '修改时间';
COMMENT ON COLUMN "group_15"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "group_2"
(
    "id"          int8 NOT NULL,
    "gid"         varchar(32),
    "name"        varchar(64),
    "username"    varchar(256),
    "sort_order"  int4,
    "create_time" timestamp,
    "update_time" timestamp,
    "del_flag"    int2,
    CONSTRAINT "_copy_65" PRIMARY KEY ("id")
);
CREATE INDEX "idx_username_copy_8" ON "group_2" USING btree (
                                                               "username" ASC
    );
COMMENT ON COLUMN "group_2"."id" IS 'ID';
COMMENT ON COLUMN "group_2"."gid" IS '分组标识';
COMMENT ON COLUMN "group_2"."name" IS '分组名称';
COMMENT ON COLUMN "group_2"."username" IS '创建分组用户名';
COMMENT ON COLUMN "group_2"."sort_order" IS '分组排序';
COMMENT ON COLUMN "group_2"."create_time" IS '创建时间';
COMMENT ON COLUMN "group_2"."update_time" IS '修改时间';
COMMENT ON COLUMN "group_2"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "group_3"
(
    "id"          int8 NOT NULL,
    "gid"         varchar(32),
    "name"        varchar(64),
    "username"    varchar(256),
    "sort_order"  int4,
    "create_time" timestamp,
    "update_time" timestamp,
    "del_flag"    int2,
    CONSTRAINT "_copy_64" PRIMARY KEY ("id")
);
CREATE INDEX "idx_username_copy_7" ON "group_3" USING btree (
                                                               "username" ASC
    );
COMMENT ON COLUMN "group_3"."id" IS 'ID';
COMMENT ON COLUMN "group_3"."gid" IS '分组标识';
COMMENT ON COLUMN "group_3"."name" IS '分组名称';
COMMENT ON COLUMN "group_3"."username" IS '创建分组用户名';
COMMENT ON COLUMN "group_3"."sort_order" IS '分组排序';
COMMENT ON COLUMN "group_3"."create_time" IS '创建时间';
COMMENT ON COLUMN "group_3"."update_time" IS '修改时间';
COMMENT ON COLUMN "group_3"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "group_4"
(
    "id"          int8 NOT NULL,
    "gid"         varchar(32),
    "name"        varchar(64),
    "username"    varchar(256),
    "sort_order"  int4,
    "create_time" timestamp,
    "update_time" timestamp,
    "del_flag"    int2,
    CONSTRAINT "_copy_63" PRIMARY KEY ("id")
);
CREATE INDEX "idx_username_copy_6" ON "group_4" USING btree (
                                                               "username" ASC
    );
COMMENT ON COLUMN "group_4"."id" IS 'ID';
COMMENT ON COLUMN "group_4"."gid" IS '分组标识';
COMMENT ON COLUMN "group_4"."name" IS '分组名称';
COMMENT ON COLUMN "group_4"."username" IS '创建分组用户名';
COMMENT ON COLUMN "group_4"."sort_order" IS '分组排序';
COMMENT ON COLUMN "group_4"."create_time" IS '创建时间';
COMMENT ON COLUMN "group_4"."update_time" IS '修改时间';
COMMENT ON COLUMN "group_4"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "group_5"
(
    "id"          int8 NOT NULL,
    "gid"         varchar(32),
    "name"        varchar(64),
    "username"    varchar(256),
    "sort_order"  int4,
    "create_time" timestamp,
    "update_time" timestamp,
    "del_flag"    int2,
    CONSTRAINT "_copy_62" PRIMARY KEY ("id")
);
CREATE INDEX "idx_username_copy_5" ON "group_5" USING btree (
                                                               "username" ASC
    );
COMMENT ON COLUMN "group_5"."id" IS 'ID';
COMMENT ON COLUMN "group_5"."gid" IS '分组标识';
COMMENT ON COLUMN "group_5"."name" IS '分组名称';
COMMENT ON COLUMN "group_5"."username" IS '创建分组用户名';
COMMENT ON COLUMN "group_5"."sort_order" IS '分组排序';
COMMENT ON COLUMN "group_5"."create_time" IS '创建时间';
COMMENT ON COLUMN "group_5"."update_time" IS '修改时间';
COMMENT ON COLUMN "group_5"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "group_6"
(
    "id"          int8 NOT NULL,
    "gid"         varchar(32),
    "name"        varchar(64),
    "username"    varchar(256),
    "sort_order"  int4,
    "create_time" timestamp,
    "update_time" timestamp,
    "del_flag"    int2,
    CONSTRAINT "_copy_61" PRIMARY KEY ("id")
);
CREATE INDEX "idx_username_copy_4" ON "group_6" USING btree (
                                                               "username" ASC
    );
COMMENT ON COLUMN "group_6"."id" IS 'ID';
COMMENT ON COLUMN "group_6"."gid" IS '分组标识';
COMMENT ON COLUMN "group_6"."name" IS '分组名称';
COMMENT ON COLUMN "group_6"."username" IS '创建分组用户名';
COMMENT ON COLUMN "group_6"."sort_order" IS '分组排序';
COMMENT ON COLUMN "group_6"."create_time" IS '创建时间';
COMMENT ON COLUMN "group_6"."update_time" IS '修改时间';
COMMENT ON COLUMN "group_6"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "group_7"
(
    "id"          int8 NOT NULL,
    "gid"         varchar(32),
    "name"        varchar(64),
    "username"    varchar(256),
    "sort_order"  int4,
    "create_time" timestamp,
    "update_time" timestamp,
    "del_flag"    int2,
    CONSTRAINT "_copy_60" PRIMARY KEY ("id")
);
CREATE INDEX "idx_username_copy_3" ON "group_7" USING btree (
                                                               "username" ASC
    );
COMMENT ON COLUMN "group_7"."id" IS 'ID';
COMMENT ON COLUMN "group_7"."gid" IS '分组标识';
COMMENT ON COLUMN "group_7"."name" IS '分组名称';
COMMENT ON COLUMN "group_7"."username" IS '创建分组用户名';
COMMENT ON COLUMN "group_7"."sort_order" IS '分组排序';
COMMENT ON COLUMN "group_7"."create_time" IS '创建时间';
COMMENT ON COLUMN "group_7"."update_time" IS '修改时间';
COMMENT ON COLUMN "group_7"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "group_8"
(
    "id"          int8 NOT NULL,
    "gid"         varchar(32),
    "name"        varchar(64),
    "username"    varchar(256),
    "sort_order"  int4,
    "create_time" timestamp,
    "update_time" timestamp,
    "del_flag"    int2,
    CONSTRAINT "_copy_59" PRIMARY KEY ("id")
);
CREATE INDEX "idx_username_copy_2" ON "group_8" USING btree (
                                                               "username" ASC
    );
COMMENT ON COLUMN "group_8"."id" IS 'ID';
COMMENT ON COLUMN "group_8"."gid" IS '分组标识';
COMMENT ON COLUMN "group_8"."name" IS '分组名称';
COMMENT ON COLUMN "group_8"."username" IS '创建分组用户名';
COMMENT ON COLUMN "group_8"."sort_order" IS '分组排序';
COMMENT ON COLUMN "group_8"."create_time" IS '创建时间';
COMMENT ON COLUMN "group_8"."update_time" IS '修改时间';
COMMENT ON COLUMN "group_8"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "group_9"
(
    "id"          int8 NOT NULL,
    "gid"         varchar(32),
    "name"        varchar(64),
    "username"    varchar(256),
    "sort_order"  int4,
    "create_time" timestamp,
    "update_time" timestamp,
    "del_flag"    int2,
    CONSTRAINT "_copy_58" PRIMARY KEY ("id")
);
CREATE INDEX "idx_username_copy_1" ON "group_9" USING btree (
                                                               "username" ASC
    );
COMMENT ON COLUMN "group_9"."id" IS 'ID';
COMMENT ON COLUMN "group_9"."gid" IS '分组标识';
COMMENT ON COLUMN "group_9"."name" IS '分组名称';
COMMENT ON COLUMN "group_9"."username" IS '创建分组用户名';
COMMENT ON COLUMN "group_9"."sort_order" IS '分组排序';
COMMENT ON COLUMN "group_9"."create_time" IS '创建时间';
COMMENT ON COLUMN "group_9"."update_time" IS '修改时间';
COMMENT ON COLUMN "group_9"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "group_unique"
(
    "id"  int8 NOT NULL,
    "gid" varchar(32),
    CONSTRAINT "_copy_57" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_gid" ON "group_unique" USING btree (
                                                                      "gid" ASC
    );
COMMENT ON COLUMN "group_unique"."id" IS 'ID';
COMMENT ON COLUMN "group_unique"."gid" IS '分组标识';

CREATE TABLE "link_0"
(
    "id"              int8 NOT NULL,
    "domain"          varchar(128),
    "short_uri"       varchar(8),
    "full_short_url"  varchar(128),
    "origin_url"      varchar(1024),
    "click_num"       int4,
    "gid"             varchar(32),
    "favicon"         varchar(256),
    "enable_status"   int2,
    "created_type"    int2,
    "valid_date_type" int2,
    "valid_date"      timestamp,
    "describe"        varchar(1024),
    "total_pv"        int4,
    "total_uv"        int4,
    "total_uip"       int4,
    "create_time"     timestamp,
    "update_time"     timestamp,
    "del_time"        int8,
    "del_flag"        int2,
    CONSTRAINT "_copy_56" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_full-short-url" ON "link_0" USING btree (
                                                                           "full_short_url" ASC,
                                                                           "del_time" ASC
    );
COMMENT ON COLUMN "link_0"."id" IS 'ID';
COMMENT ON COLUMN "link_0"."domain" IS '域名';
COMMENT ON COLUMN "link_0"."short_uri" IS '短链接';
COMMENT ON COLUMN "link_0"."full_short_url" IS '完整短链接';
COMMENT ON COLUMN "link_0"."origin_url" IS '原始链接';
COMMENT ON COLUMN "link_0"."click_num" IS '点击量';
COMMENT ON COLUMN "link_0"."gid" IS '分组标识';
COMMENT ON COLUMN "link_0"."favicon" IS '网站图标';
COMMENT ON COLUMN "link_0"."enable_status" IS '启用标识 0：启用 1：未启用';
COMMENT ON COLUMN "link_0"."created_type" IS '创建类型 0：接口创建 1：控制台创建';
COMMENT ON COLUMN "link_0"."valid_date_type" IS '有效期类型 0：永久有效 1：自定义';
COMMENT ON COLUMN "link_0"."valid_date" IS '有效期';
COMMENT ON COLUMN "link_0"."describe" IS '描述';
COMMENT ON COLUMN "link_0"."total_pv" IS '历史PV';
COMMENT ON COLUMN "link_0"."total_uv" IS '历史UV';
COMMENT ON COLUMN "link_0"."total_uip" IS '历史UIP';
COMMENT ON COLUMN "link_0"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_0"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_0"."del_time" IS '删除时间戳';
COMMENT ON COLUMN "link_0"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "link_1"
(
    "id"              int8 NOT NULL,
    "domain"          varchar(128),
    "short_uri"       varchar(8),
    "full_short_url"  varchar(128),
    "origin_url"      varchar(1024),
    "click_num"       int4,
    "gid"             varchar(32),
    "favicon"         varchar(256),
    "enable_status"   int2,
    "created_type"    int2,
    "valid_date_type" int2,
    "valid_date"      timestamp,
    "describe"        varchar(1024),
    "total_pv"        int4,
    "total_uv"        int4,
    "total_uip"       int4,
    "create_time"     timestamp,
    "update_time"     timestamp,
    "del_time"        int8,
    "del_flag"        int2,
    CONSTRAINT "_copy_55" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_full-short-url_copy_15" ON "link_1" USING btree (
                                                                                   "full_short_url" ASC,
                                                                                   "del_time" ASC
    );
COMMENT ON COLUMN "link_1"."id" IS 'ID';
COMMENT ON COLUMN "link_1"."domain" IS '域名';
COMMENT ON COLUMN "link_1"."short_uri" IS '短链接';
COMMENT ON COLUMN "link_1"."full_short_url" IS '完整短链接';
COMMENT ON COLUMN "link_1"."origin_url" IS '原始链接';
COMMENT ON COLUMN "link_1"."click_num" IS '点击量';
COMMENT ON COLUMN "link_1"."gid" IS '分组标识';
COMMENT ON COLUMN "link_1"."favicon" IS '网站图标';
COMMENT ON COLUMN "link_1"."enable_status" IS '启用标识 0：启用 1：未启用';
COMMENT ON COLUMN "link_1"."created_type" IS '创建类型 0：接口创建 1：控制台创建';
COMMENT ON COLUMN "link_1"."valid_date_type" IS '有效期类型 0：永久有效 1：自定义';
COMMENT ON COLUMN "link_1"."valid_date" IS '有效期';
COMMENT ON COLUMN "link_1"."describe" IS '描述';
COMMENT ON COLUMN "link_1"."total_pv" IS '历史PV';
COMMENT ON COLUMN "link_1"."total_uv" IS '历史UV';
COMMENT ON COLUMN "link_1"."total_uip" IS '历史UIP';
COMMENT ON COLUMN "link_1"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_1"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_1"."del_time" IS '删除时间戳';
COMMENT ON COLUMN "link_1"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "link_10"
(
    "id"              int8 NOT NULL,
    "domain"          varchar(128),
    "short_uri"       varchar(8),
    "full_short_url"  varchar(128),
    "origin_url"      varchar(1024),
    "click_num"       int4,
    "gid"             varchar(32),
    "favicon"         varchar(256),
    "enable_status"   int2,
    "created_type"    int2,
    "valid_date_type" int2,
    "valid_date"      timestamp,
    "describe"        varchar(1024),
    "total_pv"        int4,
    "total_uv"        int4,
    "total_uip"       int4,
    "create_time"     timestamp,
    "update_time"     timestamp,
    "del_time"        int8,
    "del_flag"        int2,
    CONSTRAINT "_copy_54" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_full-short-url_copy_14" ON "link_10" USING btree (
                                                                                    "full_short_url" ASC,
                                                                                    "del_time" ASC
    );
COMMENT ON COLUMN "link_10"."id" IS 'ID';
COMMENT ON COLUMN "link_10"."domain" IS '域名';
COMMENT ON COLUMN "link_10"."short_uri" IS '短链接';
COMMENT ON COLUMN "link_10"."full_short_url" IS '完整短链接';
COMMENT ON COLUMN "link_10"."origin_url" IS '原始链接';
COMMENT ON COLUMN "link_10"."click_num" IS '点击量';
COMMENT ON COLUMN "link_10"."gid" IS '分组标识';
COMMENT ON COLUMN "link_10"."favicon" IS '网站图标';
COMMENT ON COLUMN "link_10"."enable_status" IS '启用标识 0：启用 1：未启用';
COMMENT ON COLUMN "link_10"."created_type" IS '创建类型 0：接口创建 1：控制台创建';
COMMENT ON COLUMN "link_10"."valid_date_type" IS '有效期类型 0：永久有效 1：自定义';
COMMENT ON COLUMN "link_10"."valid_date" IS '有效期';
COMMENT ON COLUMN "link_10"."describe" IS '描述';
COMMENT ON COLUMN "link_10"."total_pv" IS '历史PV';
COMMENT ON COLUMN "link_10"."total_uv" IS '历史UV';
COMMENT ON COLUMN "link_10"."total_uip" IS '历史UIP';
COMMENT ON COLUMN "link_10"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_10"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_10"."del_time" IS '删除时间戳';
COMMENT ON COLUMN "link_10"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "link_11"
(
    "id"              int8 NOT NULL,
    "domain"          varchar(128),
    "short_uri"       varchar(8),
    "full_short_url"  varchar(128),
    "origin_url"      varchar(1024),
    "click_num"       int4,
    "gid"             varchar(32),
    "favicon"         varchar(256),
    "enable_status"   int2,
    "created_type"    int2,
    "valid_date_type" int2,
    "valid_date"      timestamp,
    "describe"        varchar(1024),
    "total_pv"        int4,
    "total_uv"        int4,
    "total_uip"       int4,
    "create_time"     timestamp,
    "update_time"     timestamp,
    "del_time"        int8,
    "del_flag"        int2,
    CONSTRAINT "_copy_53" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_full-short-url_copy_13" ON "link_11" USING btree (
                                                                                    "full_short_url" ASC,
                                                                                    "del_time" ASC
    );
COMMENT ON COLUMN "link_11"."id" IS 'ID';
COMMENT ON COLUMN "link_11"."domain" IS '域名';
COMMENT ON COLUMN "link_11"."short_uri" IS '短链接';
COMMENT ON COLUMN "link_11"."full_short_url" IS '完整短链接';
COMMENT ON COLUMN "link_11"."origin_url" IS '原始链接';
COMMENT ON COLUMN "link_11"."click_num" IS '点击量';
COMMENT ON COLUMN "link_11"."gid" IS '分组标识';
COMMENT ON COLUMN "link_11"."favicon" IS '网站图标';
COMMENT ON COLUMN "link_11"."enable_status" IS '启用标识 0：启用 1：未启用';
COMMENT ON COLUMN "link_11"."created_type" IS '创建类型 0：接口创建 1：控制台创建';
COMMENT ON COLUMN "link_11"."valid_date_type" IS '有效期类型 0：永久有效 1：自定义';
COMMENT ON COLUMN "link_11"."valid_date" IS '有效期';
COMMENT ON COLUMN "link_11"."describe" IS '描述';
COMMENT ON COLUMN "link_11"."total_pv" IS '历史PV';
COMMENT ON COLUMN "link_11"."total_uv" IS '历史UV';
COMMENT ON COLUMN "link_11"."total_uip" IS '历史UIP';
COMMENT ON COLUMN "link_11"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_11"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_11"."del_time" IS '删除时间戳';
COMMENT ON COLUMN "link_11"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "link_12"
(
    "id"              int8 NOT NULL,
    "domain"          varchar(128),
    "short_uri"       varchar(8),
    "full_short_url"  varchar(128),
    "origin_url"      varchar(1024),
    "click_num"       int4,
    "gid"             varchar(32),
    "favicon"         varchar(256),
    "enable_status"   int2,
    "created_type"    int2,
    "valid_date_type" int2,
    "valid_date"      timestamp,
    "describe"        varchar(1024),
    "total_pv"        int4,
    "total_uv"        int4,
    "total_uip"       int4,
    "create_time"     timestamp,
    "update_time"     timestamp,
    "del_time"        int8,
    "del_flag"        int2,
    CONSTRAINT "_copy_52" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_full-short-url_copy_12" ON "link_12" USING btree (
                                                                                    "full_short_url" ASC,
                                                                                    "del_time" ASC
    );
COMMENT ON COLUMN "link_12"."id" IS 'ID';
COMMENT ON COLUMN "link_12"."domain" IS '域名';
COMMENT ON COLUMN "link_12"."short_uri" IS '短链接';
COMMENT ON COLUMN "link_12"."full_short_url" IS '完整短链接';
COMMENT ON COLUMN "link_12"."origin_url" IS '原始链接';
COMMENT ON COLUMN "link_12"."click_num" IS '点击量';
COMMENT ON COLUMN "link_12"."gid" IS '分组标识';
COMMENT ON COLUMN "link_12"."favicon" IS '网站图标';
COMMENT ON COLUMN "link_12"."enable_status" IS '启用标识 0：启用 1：未启用';
COMMENT ON COLUMN "link_12"."created_type" IS '创建类型 0：接口创建 1：控制台创建';
COMMENT ON COLUMN "link_12"."valid_date_type" IS '有效期类型 0：永久有效 1：自定义';
COMMENT ON COLUMN "link_12"."valid_date" IS '有效期';
COMMENT ON COLUMN "link_12"."describe" IS '描述';
COMMENT ON COLUMN "link_12"."total_pv" IS '历史PV';
COMMENT ON COLUMN "link_12"."total_uv" IS '历史UV';
COMMENT ON COLUMN "link_12"."total_uip" IS '历史UIP';
COMMENT ON COLUMN "link_12"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_12"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_12"."del_time" IS '删除时间戳';
COMMENT ON COLUMN "link_12"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "link_13"
(
    "id"              int8 NOT NULL,
    "domain"          varchar(128),
    "short_uri"       varchar(8),
    "full_short_url"  varchar(128),
    "origin_url"      varchar(1024),
    "click_num"       int4,
    "gid"             varchar(32),
    "favicon"         varchar(256),
    "enable_status"   int2,
    "created_type"    int2,
    "valid_date_type" int2,
    "valid_date"      timestamp,
    "describe"        varchar(1024),
    "total_pv"        int4,
    "total_uv"        int4,
    "total_uip"       int4,
    "create_time"     timestamp,
    "update_time"     timestamp,
    "del_time"        int8,
    "del_flag"        int2,
    CONSTRAINT "_copy_51" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_full-short-url_copy_11" ON "link_13" USING btree (
                                                                                    "full_short_url" ASC,
                                                                                    "del_time" ASC
    );
COMMENT ON COLUMN "link_13"."id" IS 'ID';
COMMENT ON COLUMN "link_13"."domain" IS '域名';
COMMENT ON COLUMN "link_13"."short_uri" IS '短链接';
COMMENT ON COLUMN "link_13"."full_short_url" IS '完整短链接';
COMMENT ON COLUMN "link_13"."origin_url" IS '原始链接';
COMMENT ON COLUMN "link_13"."click_num" IS '点击量';
COMMENT ON COLUMN "link_13"."gid" IS '分组标识';
COMMENT ON COLUMN "link_13"."favicon" IS '网站图标';
COMMENT ON COLUMN "link_13"."enable_status" IS '启用标识 0：启用 1：未启用';
COMMENT ON COLUMN "link_13"."created_type" IS '创建类型 0：接口创建 1：控制台创建';
COMMENT ON COLUMN "link_13"."valid_date_type" IS '有效期类型 0：永久有效 1：自定义';
COMMENT ON COLUMN "link_13"."valid_date" IS '有效期';
COMMENT ON COLUMN "link_13"."describe" IS '描述';
COMMENT ON COLUMN "link_13"."total_pv" IS '历史PV';
COMMENT ON COLUMN "link_13"."total_uv" IS '历史UV';
COMMENT ON COLUMN "link_13"."total_uip" IS '历史UIP';
COMMENT ON COLUMN "link_13"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_13"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_13"."del_time" IS '删除时间戳';
COMMENT ON COLUMN "link_13"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "link_14"
(
    "id"              int8 NOT NULL,
    "domain"          varchar(128),
    "short_uri"       varchar(8),
    "full_short_url"  varchar(128),
    "origin_url"      varchar(1024),
    "click_num"       int4,
    "gid"             varchar(32),
    "favicon"         varchar(256),
    "enable_status"   int2,
    "created_type"    int2,
    "valid_date_type" int2,
    "valid_date"      timestamp,
    "describe"        varchar(1024),
    "total_pv"        int4,
    "total_uv"        int4,
    "total_uip"       int4,
    "create_time"     timestamp,
    "update_time"     timestamp,
    "del_time"        int8,
    "del_flag"        int2,
    CONSTRAINT "_copy_50" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_full-short-url_copy_10" ON "link_14" USING btree (
                                                                                    "full_short_url" ASC,
                                                                                    "del_time" ASC
    );
COMMENT ON COLUMN "link_14"."id" IS 'ID';
COMMENT ON COLUMN "link_14"."domain" IS '域名';
COMMENT ON COLUMN "link_14"."short_uri" IS '短链接';
COMMENT ON COLUMN "link_14"."full_short_url" IS '完整短链接';
COMMENT ON COLUMN "link_14"."origin_url" IS '原始链接';
COMMENT ON COLUMN "link_14"."click_num" IS '点击量';
COMMENT ON COLUMN "link_14"."gid" IS '分组标识';
COMMENT ON COLUMN "link_14"."favicon" IS '网站图标';
COMMENT ON COLUMN "link_14"."enable_status" IS '启用标识 0：启用 1：未启用';
COMMENT ON COLUMN "link_14"."created_type" IS '创建类型 0：接口创建 1：控制台创建';
COMMENT ON COLUMN "link_14"."valid_date_type" IS '有效期类型 0：永久有效 1：自定义';
COMMENT ON COLUMN "link_14"."valid_date" IS '有效期';
COMMENT ON COLUMN "link_14"."describe" IS '描述';
COMMENT ON COLUMN "link_14"."total_pv" IS '历史PV';
COMMENT ON COLUMN "link_14"."total_uv" IS '历史UV';
COMMENT ON COLUMN "link_14"."total_uip" IS '历史UIP';
COMMENT ON COLUMN "link_14"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_14"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_14"."del_time" IS '删除时间戳';
COMMENT ON COLUMN "link_14"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "link_15"
(
    "id"              int8 NOT NULL,
    "domain"          varchar(128),
    "short_uri"       varchar(8),
    "full_short_url"  varchar(128),
    "origin_url"      varchar(1024),
    "click_num"       int4,
    "gid"             varchar(32),
    "favicon"         varchar(256),
    "enable_status"   int2,
    "created_type"    int2,
    "valid_date_type" int2,
    "valid_date"      timestamp,
    "describe"        varchar(1024),
    "total_pv"        int4,
    "total_uv"        int4,
    "total_uip"       int4,
    "create_time"     timestamp,
    "update_time"     timestamp,
    "del_time"        int8,
    "del_flag"        int2,
    CONSTRAINT "_copy_49" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_full-short-url_copy_9" ON "link_15" USING btree (
                                                                                   "full_short_url" ASC,
                                                                                   "del_time" ASC
    );
COMMENT ON COLUMN "link_15"."id" IS 'ID';
COMMENT ON COLUMN "link_15"."domain" IS '域名';
COMMENT ON COLUMN "link_15"."short_uri" IS '短链接';
COMMENT ON COLUMN "link_15"."full_short_url" IS '完整短链接';
COMMENT ON COLUMN "link_15"."origin_url" IS '原始链接';
COMMENT ON COLUMN "link_15"."click_num" IS '点击量';
COMMENT ON COLUMN "link_15"."gid" IS '分组标识';
COMMENT ON COLUMN "link_15"."favicon" IS '网站图标';
COMMENT ON COLUMN "link_15"."enable_status" IS '启用标识 0：启用 1：未启用';
COMMENT ON COLUMN "link_15"."created_type" IS '创建类型 0：接口创建 1：控制台创建';
COMMENT ON COLUMN "link_15"."valid_date_type" IS '有效期类型 0：永久有效 1：自定义';
COMMENT ON COLUMN "link_15"."valid_date" IS '有效期';
COMMENT ON COLUMN "link_15"."describe" IS '描述';
COMMENT ON COLUMN "link_15"."total_pv" IS '历史PV';
COMMENT ON COLUMN "link_15"."total_uv" IS '历史UV';
COMMENT ON COLUMN "link_15"."total_uip" IS '历史UIP';
COMMENT ON COLUMN "link_15"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_15"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_15"."del_time" IS '删除时间戳';
COMMENT ON COLUMN "link_15"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "link_2"
(
    "id"              int8 NOT NULL,
    "domain"          varchar(128),
    "short_uri"       varchar(8),
    "full_short_url"  varchar(128),
    "origin_url"      varchar(1024),
    "click_num"       int4,
    "gid"             varchar(32),
    "favicon"         varchar(256),
    "enable_status"   int2,
    "created_type"    int2,
    "valid_date_type" int2,
    "valid_date"      timestamp,
    "describe"        varchar(1024),
    "total_pv"        int4,
    "total_uv"        int4,
    "total_uip"       int4,
    "create_time"     timestamp,
    "update_time"     timestamp,
    "del_time"        int8,
    "del_flag"        int2,
    CONSTRAINT "_copy_48" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_full-short-url_copy_8" ON "link_2" USING btree (
                                                                                  "full_short_url" ASC,
                                                                                  "del_time" ASC
    );
COMMENT ON COLUMN "link_2"."id" IS 'ID';
COMMENT ON COLUMN "link_2"."domain" IS '域名';
COMMENT ON COLUMN "link_2"."short_uri" IS '短链接';
COMMENT ON COLUMN "link_2"."full_short_url" IS '完整短链接';
COMMENT ON COLUMN "link_2"."origin_url" IS '原始链接';
COMMENT ON COLUMN "link_2"."click_num" IS '点击量';
COMMENT ON COLUMN "link_2"."gid" IS '分组标识';
COMMENT ON COLUMN "link_2"."favicon" IS '网站图标';
COMMENT ON COLUMN "link_2"."enable_status" IS '启用标识 0：启用 1：未启用';
COMMENT ON COLUMN "link_2"."created_type" IS '创建类型 0：接口创建 1：控制台创建';
COMMENT ON COLUMN "link_2"."valid_date_type" IS '有效期类型 0：永久有效 1：自定义';
COMMENT ON COLUMN "link_2"."valid_date" IS '有效期';
COMMENT ON COLUMN "link_2"."describe" IS '描述';
COMMENT ON COLUMN "link_2"."total_pv" IS '历史PV';
COMMENT ON COLUMN "link_2"."total_uv" IS '历史UV';
COMMENT ON COLUMN "link_2"."total_uip" IS '历史UIP';
COMMENT ON COLUMN "link_2"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_2"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_2"."del_time" IS '删除时间戳';
COMMENT ON COLUMN "link_2"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "link_3"
(
    "id"              int8 NOT NULL,
    "domain"          varchar(128),
    "short_uri"       varchar(8),
    "full_short_url"  varchar(128),
    "origin_url"      varchar(1024),
    "click_num"       int4,
    "gid"             varchar(32),
    "favicon"         varchar(256),
    "enable_status"   int2,
    "created_type"    int2,
    "valid_date_type" int2,
    "valid_date"      timestamp,
    "describe"        varchar(1024),
    "total_pv"        int4,
    "total_uv"        int4,
    "total_uip"       int4,
    "create_time"     timestamp,
    "update_time"     timestamp,
    "del_time"        int8,
    "del_flag"        int2,
    CONSTRAINT "_copy_47" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_full-short-url_copy_7" ON "link_3" USING btree (
                                                                                  "full_short_url" ASC,
                                                                                  "del_time" ASC
    );
COMMENT ON COLUMN "link_3"."id" IS 'ID';
COMMENT ON COLUMN "link_3"."domain" IS '域名';
COMMENT ON COLUMN "link_3"."short_uri" IS '短链接';
COMMENT ON COLUMN "link_3"."full_short_url" IS '完整短链接';
COMMENT ON COLUMN "link_3"."origin_url" IS '原始链接';
COMMENT ON COLUMN "link_3"."click_num" IS '点击量';
COMMENT ON COLUMN "link_3"."gid" IS '分组标识';
COMMENT ON COLUMN "link_3"."favicon" IS '网站图标';
COMMENT ON COLUMN "link_3"."enable_status" IS '启用标识 0：启用 1：未启用';
COMMENT ON COLUMN "link_3"."created_type" IS '创建类型 0：接口创建 1：控制台创建';
COMMENT ON COLUMN "link_3"."valid_date_type" IS '有效期类型 0：永久有效 1：自定义';
COMMENT ON COLUMN "link_3"."valid_date" IS '有效期';
COMMENT ON COLUMN "link_3"."describe" IS '描述';
COMMENT ON COLUMN "link_3"."total_pv" IS '历史PV';
COMMENT ON COLUMN "link_3"."total_uv" IS '历史UV';
COMMENT ON COLUMN "link_3"."total_uip" IS '历史UIP';
COMMENT ON COLUMN "link_3"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_3"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_3"."del_time" IS '删除时间戳';
COMMENT ON COLUMN "link_3"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "link_4"
(
    "id"              int8 NOT NULL,
    "domain"          varchar(128),
    "short_uri"       varchar(8),
    "full_short_url"  varchar(128),
    "origin_url"      varchar(1024),
    "click_num"       int4,
    "gid"             varchar(32),
    "favicon"         varchar(256),
    "enable_status"   int2,
    "created_type"    int2,
    "valid_date_type" int2,
    "valid_date"      timestamp,
    "describe"        varchar(1024),
    "total_pv"        int4,
    "total_uv"        int4,
    "total_uip"       int4,
    "create_time"     timestamp,
    "update_time"     timestamp,
    "del_time"        int8,
    "del_flag"        int2,
    CONSTRAINT "_copy_46" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_full-short-url_copy_6" ON "link_4" USING btree (
                                                                                  "full_short_url" ASC,
                                                                                  "del_time" ASC
    );
COMMENT ON COLUMN "link_4"."id" IS 'ID';
COMMENT ON COLUMN "link_4"."domain" IS '域名';
COMMENT ON COLUMN "link_4"."short_uri" IS '短链接';
COMMENT ON COLUMN "link_4"."full_short_url" IS '完整短链接';
COMMENT ON COLUMN "link_4"."origin_url" IS '原始链接';
COMMENT ON COLUMN "link_4"."click_num" IS '点击量';
COMMENT ON COLUMN "link_4"."gid" IS '分组标识';
COMMENT ON COLUMN "link_4"."favicon" IS '网站图标';
COMMENT ON COLUMN "link_4"."enable_status" IS '启用标识 0：启用 1：未启用';
COMMENT ON COLUMN "link_4"."created_type" IS '创建类型 0：接口创建 1：控制台创建';
COMMENT ON COLUMN "link_4"."valid_date_type" IS '有效期类型 0：永久有效 1：自定义';
COMMENT ON COLUMN "link_4"."valid_date" IS '有效期';
COMMENT ON COLUMN "link_4"."describe" IS '描述';
COMMENT ON COLUMN "link_4"."total_pv" IS '历史PV';
COMMENT ON COLUMN "link_4"."total_uv" IS '历史UV';
COMMENT ON COLUMN "link_4"."total_uip" IS '历史UIP';
COMMENT ON COLUMN "link_4"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_4"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_4"."del_time" IS '删除时间戳';
COMMENT ON COLUMN "link_4"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "link_5"
(
    "id"              int8 NOT NULL,
    "domain"          varchar(128),
    "short_uri"       varchar(8),
    "full_short_url"  varchar(128),
    "origin_url"      varchar(1024),
    "click_num"       int4,
    "gid"             varchar(32),
    "favicon"         varchar(256),
    "enable_status"   int2,
    "created_type"    int2,
    "valid_date_type" int2,
    "valid_date"      timestamp,
    "describe"        varchar(1024),
    "total_pv"        int4,
    "total_uv"        int4,
    "total_uip"       int4,
    "create_time"     timestamp,
    "update_time"     timestamp,
    "del_time"        int8,
    "del_flag"        int2,
    CONSTRAINT "_copy_45" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_full-short-url_copy_5" ON "link_5" USING btree (
                                                                                  "full_short_url" ASC,
                                                                                  "del_time" ASC
    );
COMMENT ON COLUMN "link_5"."id" IS 'ID';
COMMENT ON COLUMN "link_5"."domain" IS '域名';
COMMENT ON COLUMN "link_5"."short_uri" IS '短链接';
COMMENT ON COLUMN "link_5"."full_short_url" IS '完整短链接';
COMMENT ON COLUMN "link_5"."origin_url" IS '原始链接';
COMMENT ON COLUMN "link_5"."click_num" IS '点击量';
COMMENT ON COLUMN "link_5"."gid" IS '分组标识';
COMMENT ON COLUMN "link_5"."favicon" IS '网站图标';
COMMENT ON COLUMN "link_5"."enable_status" IS '启用标识 0：启用 1：未启用';
COMMENT ON COLUMN "link_5"."created_type" IS '创建类型 0：接口创建 1：控制台创建';
COMMENT ON COLUMN "link_5"."valid_date_type" IS '有效期类型 0：永久有效 1：自定义';
COMMENT ON COLUMN "link_5"."valid_date" IS '有效期';
COMMENT ON COLUMN "link_5"."describe" IS '描述';
COMMENT ON COLUMN "link_5"."total_pv" IS '历史PV';
COMMENT ON COLUMN "link_5"."total_uv" IS '历史UV';
COMMENT ON COLUMN "link_5"."total_uip" IS '历史UIP';
COMMENT ON COLUMN "link_5"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_5"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_5"."del_time" IS '删除时间戳';
COMMENT ON COLUMN "link_5"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "link_6"
(
    "id"              int8 NOT NULL,
    "domain"          varchar(128),
    "short_uri"       varchar(8),
    "full_short_url"  varchar(128),
    "origin_url"      varchar(1024),
    "click_num"       int4,
    "gid"             varchar(32),
    "favicon"         varchar(256),
    "enable_status"   int2,
    "created_type"    int2,
    "valid_date_type" int2,
    "valid_date"      timestamp,
    "describe"        varchar(1024),
    "total_pv"        int4,
    "total_uv"        int4,
    "total_uip"       int4,
    "create_time"     timestamp,
    "update_time"     timestamp,
    "del_time"        int8,
    "del_flag"        int2,
    CONSTRAINT "_copy_44" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_full-short-url_copy_4" ON "link_6" USING btree (
                                                                                  "full_short_url" ASC,
                                                                                  "del_time" ASC
    );
COMMENT ON COLUMN "link_6"."id" IS 'ID';
COMMENT ON COLUMN "link_6"."domain" IS '域名';
COMMENT ON COLUMN "link_6"."short_uri" IS '短链接';
COMMENT ON COLUMN "link_6"."full_short_url" IS '完整短链接';
COMMENT ON COLUMN "link_6"."origin_url" IS '原始链接';
COMMENT ON COLUMN "link_6"."click_num" IS '点击量';
COMMENT ON COLUMN "link_6"."gid" IS '分组标识';
COMMENT ON COLUMN "link_6"."favicon" IS '网站图标';
COMMENT ON COLUMN "link_6"."enable_status" IS '启用标识 0：启用 1：未启用';
COMMENT ON COLUMN "link_6"."created_type" IS '创建类型 0：接口创建 1：控制台创建';
COMMENT ON COLUMN "link_6"."valid_date_type" IS '有效期类型 0：永久有效 1：自定义';
COMMENT ON COLUMN "link_6"."valid_date" IS '有效期';
COMMENT ON COLUMN "link_6"."describe" IS '描述';
COMMENT ON COLUMN "link_6"."total_pv" IS '历史PV';
COMMENT ON COLUMN "link_6"."total_uv" IS '历史UV';
COMMENT ON COLUMN "link_6"."total_uip" IS '历史UIP';
COMMENT ON COLUMN "link_6"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_6"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_6"."del_time" IS '删除时间戳';
COMMENT ON COLUMN "link_6"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "link_7"
(
    "id"              int8 NOT NULL,
    "domain"          varchar(128),
    "short_uri"       varchar(8),
    "full_short_url"  varchar(128),
    "origin_url"      varchar(1024),
    "click_num"       int4,
    "gid"             varchar(32),
    "favicon"         varchar(256),
    "enable_status"   int2,
    "created_type"    int2,
    "valid_date_type" int2,
    "valid_date"      timestamp,
    "describe"        varchar(1024),
    "total_pv"        int4,
    "total_uv"        int4,
    "total_uip"       int4,
    "create_time"     timestamp,
    "update_time"     timestamp,
    "del_time"        int8,
    "del_flag"        int2,
    CONSTRAINT "_copy_43" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_full-short-url_copy_3" ON "link_7" USING btree (
                                                                                  "full_short_url" ASC,
                                                                                  "del_time" ASC
    );
COMMENT ON COLUMN "link_7"."id" IS 'ID';
COMMENT ON COLUMN "link_7"."domain" IS '域名';
COMMENT ON COLUMN "link_7"."short_uri" IS '短链接';
COMMENT ON COLUMN "link_7"."full_short_url" IS '完整短链接';
COMMENT ON COLUMN "link_7"."origin_url" IS '原始链接';
COMMENT ON COLUMN "link_7"."click_num" IS '点击量';
COMMENT ON COLUMN "link_7"."gid" IS '分组标识';
COMMENT ON COLUMN "link_7"."favicon" IS '网站图标';
COMMENT ON COLUMN "link_7"."enable_status" IS '启用标识 0：启用 1：未启用';
COMMENT ON COLUMN "link_7"."created_type" IS '创建类型 0：接口创建 1：控制台创建';
COMMENT ON COLUMN "link_7"."valid_date_type" IS '有效期类型 0：永久有效 1：自定义';
COMMENT ON COLUMN "link_7"."valid_date" IS '有效期';
COMMENT ON COLUMN "link_7"."describe" IS '描述';
COMMENT ON COLUMN "link_7"."total_pv" IS '历史PV';
COMMENT ON COLUMN "link_7"."total_uv" IS '历史UV';
COMMENT ON COLUMN "link_7"."total_uip" IS '历史UIP';
COMMENT ON COLUMN "link_7"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_7"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_7"."del_time" IS '删除时间戳';
COMMENT ON COLUMN "link_7"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "link_8"
(
    "id"              int8 NOT NULL,
    "domain"          varchar(128),
    "short_uri"       varchar(8),
    "full_short_url"  varchar(128),
    "origin_url"      varchar(1024),
    "click_num"       int4,
    "gid"             varchar(32),
    "favicon"         varchar(256),
    "enable_status"   int2,
    "created_type"    int2,
    "valid_date_type" int2,
    "valid_date"      timestamp,
    "describe"        varchar(1024),
    "total_pv"        int4,
    "total_uv"        int4,
    "total_uip"       int4,
    "create_time"     timestamp,
    "update_time"     timestamp,
    "del_time"        int8,
    "del_flag"        int2,
    CONSTRAINT "_copy_42" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_full-short-url_copy_2" ON "link_8" USING btree (
                                                                                  "full_short_url" ASC,
                                                                                  "del_time" ASC
    );
COMMENT ON COLUMN "link_8"."id" IS 'ID';
COMMENT ON COLUMN "link_8"."domain" IS '域名';
COMMENT ON COLUMN "link_8"."short_uri" IS '短链接';
COMMENT ON COLUMN "link_8"."full_short_url" IS '完整短链接';
COMMENT ON COLUMN "link_8"."origin_url" IS '原始链接';
COMMENT ON COLUMN "link_8"."click_num" IS '点击量';
COMMENT ON COLUMN "link_8"."gid" IS '分组标识';
COMMENT ON COLUMN "link_8"."favicon" IS '网站图标';
COMMENT ON COLUMN "link_8"."enable_status" IS '启用标识 0：启用 1：未启用';
COMMENT ON COLUMN "link_8"."created_type" IS '创建类型 0：接口创建 1：控制台创建';
COMMENT ON COLUMN "link_8"."valid_date_type" IS '有效期类型 0：永久有效 1：自定义';
COMMENT ON COLUMN "link_8"."valid_date" IS '有效期';
COMMENT ON COLUMN "link_8"."describe" IS '描述';
COMMENT ON COLUMN "link_8"."total_pv" IS '历史PV';
COMMENT ON COLUMN "link_8"."total_uv" IS '历史UV';
COMMENT ON COLUMN "link_8"."total_uip" IS '历史UIP';
COMMENT ON COLUMN "link_8"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_8"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_8"."del_time" IS '删除时间戳';
COMMENT ON COLUMN "link_8"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "link_9"
(
    "id"              int8 NOT NULL,
    "domain"          varchar(128),
    "short_uri"       varchar(8),
    "full_short_url"  varchar(128),
    "origin_url"      varchar(1024),
    "click_num"       int4,
    "gid"             varchar(32),
    "favicon"         varchar(256),
    "enable_status"   int2,
    "created_type"    int2,
    "valid_date_type" int2,
    "valid_date"      timestamp,
    "describe"        varchar(1024),
    "total_pv"        int4,
    "total_uv"        int4,
    "total_uip"       int4,
    "create_time"     timestamp,
    "update_time"     timestamp,
    "del_time"        int8,
    "del_flag"        int2,
    CONSTRAINT "_copy_41" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_full-short-url_copy_1" ON "link_9" USING btree (
                                                                                  "full_short_url" ASC,
                                                                                  "del_time" ASC
    );
COMMENT ON COLUMN "link_9"."id" IS 'ID';
COMMENT ON COLUMN "link_9"."domain" IS '域名';
COMMENT ON COLUMN "link_9"."short_uri" IS '短链接';
COMMENT ON COLUMN "link_9"."full_short_url" IS '完整短链接';
COMMENT ON COLUMN "link_9"."origin_url" IS '原始链接';
COMMENT ON COLUMN "link_9"."click_num" IS '点击量';
COMMENT ON COLUMN "link_9"."gid" IS '分组标识';
COMMENT ON COLUMN "link_9"."favicon" IS '网站图标';
COMMENT ON COLUMN "link_9"."enable_status" IS '启用标识 0：启用 1：未启用';
COMMENT ON COLUMN "link_9"."created_type" IS '创建类型 0：接口创建 1：控制台创建';
COMMENT ON COLUMN "link_9"."valid_date_type" IS '有效期类型 0：永久有效 1：自定义';
COMMENT ON COLUMN "link_9"."valid_date" IS '有效期';
COMMENT ON COLUMN "link_9"."describe" IS '描述';
COMMENT ON COLUMN "link_9"."total_pv" IS '历史PV';
COMMENT ON COLUMN "link_9"."total_uv" IS '历史UV';
COMMENT ON COLUMN "link_9"."total_uip" IS '历史UIP';
COMMENT ON COLUMN "link_9"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_9"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_9"."del_time" IS '删除时间戳';
COMMENT ON COLUMN "link_9"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "link_access_logs"
(
    "id"             int8 NOT NULL,
    "full_short_url" varchar(128),
    "user"           varchar(64),
    "ip"             varchar(64),
    "browser"        varchar(64),
    "os"             varchar(64),
    "network"        varchar(64),
    "device"         varchar(64),
    "locale"         varchar(256),
    "create_time"    timestamp,
    "update_time"    timestamp,
    "del_flag"       int2,
    CONSTRAINT "_copy_40" PRIMARY KEY ("id")
);
CREATE INDEX "idx_full_short_url" ON "link_access_logs" USING btree (
                                                                       "full_short_url" ASC
    );
COMMENT ON COLUMN "link_access_logs"."id" IS 'ID';
COMMENT ON COLUMN "link_access_logs"."full_short_url" IS '完整短链接';
COMMENT ON COLUMN "link_access_logs"."user" IS '用户信息';
COMMENT ON COLUMN "link_access_logs"."ip" IS 'IP';
COMMENT ON COLUMN "link_access_logs"."browser" IS '浏览器';
COMMENT ON COLUMN "link_access_logs"."os" IS '操作系统';
COMMENT ON COLUMN "link_access_logs"."network" IS '访问网络';
COMMENT ON COLUMN "link_access_logs"."device" IS '访问设备';
COMMENT ON COLUMN "link_access_logs"."locale" IS '地区';
COMMENT ON COLUMN "link_access_logs"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_access_logs"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_access_logs"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "link_access_stats"
(
    "id"             int8 NOT NULL,
    "full_short_url" varchar(128),
    "date"           date,
    "pv"             int4,
    "uv"             int4,
    "uip"            int4,
    "hour"           int4,
    "weekday"        int4,
    "create_time"    timestamp,
    "update_time"    timestamp,
    "del_flag"       int2,
    CONSTRAINT "_copy_39" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_access_stats" ON "link_access_stats" USING btree (
                                                                                    "full_short_url" ASC,
                                                                                    "date" ASC,
                                                                                    "hour" ASC
    );
COMMENT ON COLUMN "link_access_stats"."id" IS 'ID';
COMMENT ON COLUMN "link_access_stats"."full_short_url" IS '完整短链接';
COMMENT ON COLUMN "link_access_stats"."date" IS '日期';
COMMENT ON COLUMN "link_access_stats"."pv" IS '访问量';
COMMENT ON COLUMN "link_access_stats"."uv" IS '独立访客数';
COMMENT ON COLUMN "link_access_stats"."uip" IS '独立IP数';
COMMENT ON COLUMN "link_access_stats"."hour" IS '小时';
COMMENT ON COLUMN "link_access_stats"."weekday" IS '星期';
COMMENT ON COLUMN "link_access_stats"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_access_stats"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_access_stats"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "link_browser_stats"
(
    "id"             int8 NOT NULL,
    "full_short_url" varchar(128),
    "date"           date,
    "cnt"            int4,
    "browser"        varchar(64),
    "create_time"    timestamp,
    "update_time"    timestamp,
    "del_flag"       int2,
    CONSTRAINT "_copy_38" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_browser_stats" ON "link_browser_stats" USING btree (
                                                                                      "full_short_url" ASC,
                                                                                      "date" ASC,
                                                                                      "browser" ASC
    );
COMMENT ON COLUMN "link_browser_stats"."id" IS 'ID';
COMMENT ON COLUMN "link_browser_stats"."full_short_url" IS '完整短链接';
COMMENT ON COLUMN "link_browser_stats"."date" IS '日期';
COMMENT ON COLUMN "link_browser_stats"."cnt" IS '访问量';
COMMENT ON COLUMN "link_browser_stats"."browser" IS '浏览器';
COMMENT ON COLUMN "link_browser_stats"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_browser_stats"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_browser_stats"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "link_device_stats"
(
    "id"             int8 NOT NULL,
    "full_short_url" varchar(128),
    "date"           date,
    "cnt"            int4,
    "device"         varchar(64),
    "create_time"    timestamp,
    "update_time"    timestamp,
    "del_flag"       int2,
    CONSTRAINT "_copy_37" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_browser_stats_copy_2" ON "link_device_stats" USING btree (
                                                                                            "full_short_url" ASC,
                                                                                            "date" ASC,
                                                                                            "device" ASC
    );
COMMENT ON COLUMN "link_device_stats"."id" IS 'ID';
COMMENT ON COLUMN "link_device_stats"."full_short_url" IS '完整短链接';
COMMENT ON COLUMN "link_device_stats"."date" IS '日期';
COMMENT ON COLUMN "link_device_stats"."cnt" IS '访问量';
COMMENT ON COLUMN "link_device_stats"."device" IS '访问设备';
COMMENT ON COLUMN "link_device_stats"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_device_stats"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_device_stats"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "link_goto_0"
(
    "id"             int8 NOT NULL,
    "gid"            varchar(32),
    "full_short_url" varchar(128),
    CONSTRAINT "_copy_36" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_full_short_url_copy_16" ON "link_goto_0" USING btree (
                                                                                 "full_short_url" ASC
    );
COMMENT ON COLUMN "link_goto_0"."id" IS 'ID';
COMMENT ON COLUMN "link_goto_0"."gid" IS '分组标识';
COMMENT ON COLUMN "link_goto_0"."full_short_url" IS '完整短链接';

CREATE TABLE "link_goto_1"
(
    "id"             int8 NOT NULL,
    "gid"            varchar(32),
    "full_short_url" varchar(128),
    CONSTRAINT "_copy_35" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_full_short_url_copy_15" ON "link_goto_1" USING btree (
                                                                                 "full_short_url" ASC
    );
COMMENT ON COLUMN "link_goto_1"."id" IS 'ID';
COMMENT ON COLUMN "link_goto_1"."gid" IS '分组标识';
COMMENT ON COLUMN "link_goto_1"."full_short_url" IS '完整短链接';

CREATE TABLE "link_goto_10"
(
    "id"             int8 NOT NULL,
    "gid"            varchar(32),
    "full_short_url" varchar(128),
    CONSTRAINT "_copy_34" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_full_short_url_copy_14" ON "link_goto_10" USING btree (
                                                                                  "full_short_url" ASC
    );
COMMENT ON COLUMN "link_goto_10"."id" IS 'ID';
COMMENT ON COLUMN "link_goto_10"."gid" IS '分组标识';
COMMENT ON COLUMN "link_goto_10"."full_short_url" IS '完整短链接';

CREATE TABLE "link_goto_11"
(
    "id"             int8 NOT NULL,
    "gid"            varchar(32),
    "full_short_url" varchar(128),
    CONSTRAINT "_copy_33" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_full_short_url_copy_13" ON "link_goto_11" USING btree (
                                                                                  "full_short_url" ASC
    );
COMMENT ON COLUMN "link_goto_11"."id" IS 'ID';
COMMENT ON COLUMN "link_goto_11"."gid" IS '分组标识';
COMMENT ON COLUMN "link_goto_11"."full_short_url" IS '完整短链接';

CREATE TABLE "link_goto_12"
(
    "id"             int8 NOT NULL,
    "gid"            varchar(32),
    "full_short_url" varchar(128),
    CONSTRAINT "_copy_32" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_full_short_url_copy_12" ON "link_goto_12" USING btree (
                                                                                  "full_short_url" ASC
    );
COMMENT ON COLUMN "link_goto_12"."id" IS 'ID';
COMMENT ON COLUMN "link_goto_12"."gid" IS '分组标识';
COMMENT ON COLUMN "link_goto_12"."full_short_url" IS '完整短链接';

CREATE TABLE "link_goto_13"
(
    "id"             int8 NOT NULL,
    "gid"            varchar(32),
    "full_short_url" varchar(128),
    CONSTRAINT "_copy_31" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_full_short_url_copy_11" ON "link_goto_13" USING btree (
                                                                                  "full_short_url" ASC
    );
COMMENT ON COLUMN "link_goto_13"."id" IS 'ID';
COMMENT ON COLUMN "link_goto_13"."gid" IS '分组标识';
COMMENT ON COLUMN "link_goto_13"."full_short_url" IS '完整短链接';

CREATE TABLE "link_goto_14"
(
    "id"             int8 NOT NULL,
    "gid"            varchar(32),
    "full_short_url" varchar(128),
    CONSTRAINT "_copy_30" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_full_short_url_copy_10" ON "link_goto_14" USING btree (
                                                                                  "full_short_url" ASC
    );
COMMENT ON COLUMN "link_goto_14"."id" IS 'ID';
COMMENT ON COLUMN "link_goto_14"."gid" IS '分组标识';
COMMENT ON COLUMN "link_goto_14"."full_short_url" IS '完整短链接';

CREATE TABLE "link_goto_15"
(
    "id"             int8 NOT NULL,
    "gid"            varchar(32),
    "full_short_url" varchar(128),
    CONSTRAINT "_copy_29" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_full_short_url_copy_9" ON "link_goto_15" USING btree (
                                                                                 "full_short_url" ASC
    );
COMMENT ON COLUMN "link_goto_15"."id" IS 'ID';
COMMENT ON COLUMN "link_goto_15"."gid" IS '分组标识';
COMMENT ON COLUMN "link_goto_15"."full_short_url" IS '完整短链接';

CREATE TABLE "link_goto_2"
(
    "id"             int8 NOT NULL,
    "gid"            varchar(32),
    "full_short_url" varchar(128),
    CONSTRAINT "_copy_28" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_full_short_url_copy_8" ON "link_goto_2" USING btree (
                                                                                "full_short_url" ASC
    );
COMMENT ON COLUMN "link_goto_2"."id" IS 'ID';
COMMENT ON COLUMN "link_goto_2"."gid" IS '分组标识';
COMMENT ON COLUMN "link_goto_2"."full_short_url" IS '完整短链接';

CREATE TABLE "link_goto_3"
(
    "id"             int8 NOT NULL,
    "gid"            varchar(32),
    "full_short_url" varchar(128),
    CONSTRAINT "_copy_27" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_full_short_url_copy_7" ON "link_goto_3" USING btree (
                                                                                "full_short_url" ASC
    );
COMMENT ON COLUMN "link_goto_3"."id" IS 'ID';
COMMENT ON COLUMN "link_goto_3"."gid" IS '分组标识';
COMMENT ON COLUMN "link_goto_3"."full_short_url" IS '完整短链接';

CREATE TABLE "link_goto_4"
(
    "id"             int8 NOT NULL,
    "gid"            varchar(32),
    "full_short_url" varchar(128),
    CONSTRAINT "_copy_26" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_full_short_url_copy_6" ON "link_goto_4" USING btree (
                                                                                "full_short_url" ASC
    );
COMMENT ON COLUMN "link_goto_4"."id" IS 'ID';
COMMENT ON COLUMN "link_goto_4"."gid" IS '分组标识';
COMMENT ON COLUMN "link_goto_4"."full_short_url" IS '完整短链接';

CREATE TABLE "link_goto_5"
(
    "id"             int8 NOT NULL,
    "gid"            varchar(32),
    "full_short_url" varchar(128),
    CONSTRAINT "_copy_25" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_full_short_url_copy_5" ON "link_goto_5" USING btree (
                                                                                "full_short_url" ASC
    );
COMMENT ON COLUMN "link_goto_5"."id" IS 'ID';
COMMENT ON COLUMN "link_goto_5"."gid" IS '分组标识';
COMMENT ON COLUMN "link_goto_5"."full_short_url" IS '完整短链接';

CREATE TABLE "link_goto_6"
(
    "id"             int8 NOT NULL,
    "gid"            varchar(32),
    "full_short_url" varchar(128),
    CONSTRAINT "_copy_24" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_full_short_url_copy_4" ON "link_goto_6" USING btree (
                                                                                "full_short_url" ASC
    );
COMMENT ON COLUMN "link_goto_6"."id" IS 'ID';
COMMENT ON COLUMN "link_goto_6"."gid" IS '分组标识';
COMMENT ON COLUMN "link_goto_6"."full_short_url" IS '完整短链接';

CREATE TABLE "link_goto_7"
(
    "id"             int8 NOT NULL,
    "gid"            varchar(32),
    "full_short_url" varchar(128),
    CONSTRAINT "_copy_23" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_full_short_url_copy_3" ON "link_goto_7" USING btree (
                                                                                "full_short_url" ASC
    );
COMMENT ON COLUMN "link_goto_7"."id" IS 'ID';
COMMENT ON COLUMN "link_goto_7"."gid" IS '分组标识';
COMMENT ON COLUMN "link_goto_7"."full_short_url" IS '完整短链接';

CREATE TABLE "link_goto_8"
(
    "id"             int8 NOT NULL,
    "gid"            varchar(32),
    "full_short_url" varchar(128),
    CONSTRAINT "_copy_22" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_full_short_url_copy_2" ON "link_goto_8" USING btree (
                                                                                "full_short_url" ASC
    );
COMMENT ON COLUMN "link_goto_8"."id" IS 'ID';
COMMENT ON COLUMN "link_goto_8"."gid" IS '分组标识';
COMMENT ON COLUMN "link_goto_8"."full_short_url" IS '完整短链接';

CREATE TABLE "link_goto_9"
(
    "id"             int8 NOT NULL,
    "gid"            varchar(32),
    "full_short_url" varchar(128),
    CONSTRAINT "_copy_21" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_full_short_url_copy_1" ON "link_goto_9" USING btree (
                                                                                "full_short_url" ASC
    );
COMMENT ON COLUMN "link_goto_9"."id" IS 'ID';
COMMENT ON COLUMN "link_goto_9"."gid" IS '分组标识';
COMMENT ON COLUMN "link_goto_9"."full_short_url" IS '完整短链接';

CREATE TABLE "link_locale_stats"
(
    "id"             int8 NOT NULL,
    "full_short_url" varchar(128),
    "date"           date,
    "cnt"            int4,
    "province"       varchar(64),
    "city"           varchar(64),
    "adcode"         varchar(64),
    "country"        varchar(64),
    "create_time"    timestamp,
    "update_time"    timestamp,
    "del_flag"       int2,
    CONSTRAINT "_copy_20" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_locale_stats" ON "link_locale_stats" USING btree (
                                                                                    "full_short_url" ASC,
                                                                                    "date" ASC,
                                                                                    "adcode" ASC,
                                                                                    "province" ASC
    );
COMMENT ON COLUMN "link_locale_stats"."id" IS 'ID';
COMMENT ON COLUMN "link_locale_stats"."full_short_url" IS '完整短链接';
COMMENT ON COLUMN "link_locale_stats"."date" IS '日期';
COMMENT ON COLUMN "link_locale_stats"."cnt" IS '访问量';
COMMENT ON COLUMN "link_locale_stats"."province" IS '省份名称';
COMMENT ON COLUMN "link_locale_stats"."city" IS '市名称';
COMMENT ON COLUMN "link_locale_stats"."adcode" IS '城市编码';
COMMENT ON COLUMN "link_locale_stats"."country" IS '国家标识';
COMMENT ON COLUMN "link_locale_stats"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_locale_stats"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_locale_stats"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "link_network_stats"
(
    "id"             int8 NOT NULL,
    "full_short_url" varchar(128),
    "date"           date,
    "cnt"            int4,
    "network"        varchar(64),
    "create_time"    timestamp,
    "update_time"    timestamp,
    "del_flag"       int2,
    CONSTRAINT "_copy_19" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_browser_stats_copy_1" ON "link_network_stats" USING btree (
                                                                                             "full_short_url" ASC,
                                                                                             "date" ASC,
                                                                                             "network" ASC
    );
COMMENT ON COLUMN "link_network_stats"."id" IS 'ID';
COMMENT ON COLUMN "link_network_stats"."full_short_url" IS '完整短链接';
COMMENT ON COLUMN "link_network_stats"."date" IS '日期';
COMMENT ON COLUMN "link_network_stats"."cnt" IS '访问量';
COMMENT ON COLUMN "link_network_stats"."network" IS '访问网络';
COMMENT ON COLUMN "link_network_stats"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_network_stats"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_network_stats"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "link_os_stats"
(
    "id"             int8 NOT NULL,
    "full_short_url" varchar(128),
    "date"           date,
    "cnt"            int4,
    "os"             varchar(64),
    "create_time"    timestamp,
    "update_time"    timestamp,
    "del_flag"       int2,
    CONSTRAINT "_copy_18" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_os_stats" ON "link_os_stats" USING btree (
                                                                            "full_short_url" ASC,
                                                                            "date" ASC,
                                                                            "os" ASC
    );
COMMENT ON COLUMN "link_os_stats"."id" IS 'ID';
COMMENT ON COLUMN "link_os_stats"."full_short_url" IS '完整短链接';
COMMENT ON COLUMN "link_os_stats"."date" IS '日期';
COMMENT ON COLUMN "link_os_stats"."cnt" IS '访问量';
COMMENT ON COLUMN "link_os_stats"."os" IS '操作系统';
COMMENT ON COLUMN "link_os_stats"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_os_stats"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_os_stats"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "link_stats_today"
(
    "id"             int8 NOT NULL,
    "full_short_url" varchar(128),
    "date"           date,
    "today_pv"       int4,
    "today_uv"       int4,
    "today_uip"      int4,
    "create_time"    timestamp,
    "update_time"    timestamp,
    "del_flag"       int2,
    CONSTRAINT "_copy_17" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_today_stats" ON "link_stats_today" USING btree (
                                                                                  "full_short_url" ASC,
                                                                                  "date" ASC
    );
COMMENT ON COLUMN "link_stats_today"."id" IS 'ID';
COMMENT ON COLUMN "link_stats_today"."full_short_url" IS '短链接';
COMMENT ON COLUMN "link_stats_today"."date" IS '日期';
COMMENT ON COLUMN "link_stats_today"."today_pv" IS '今日PV';
COMMENT ON COLUMN "link_stats_today"."today_uv" IS '今日UV';
COMMENT ON COLUMN "link_stats_today"."today_uip" IS '今日IP数';
COMMENT ON COLUMN "link_stats_today"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_stats_today"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_stats_today"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "user_0"
(
    "id"            int8 NOT NULL,
    "username"      varchar(256),
    "password"      varchar(512),
    "real_name"     varchar(256),
    "phone"         varchar(128),
    "mail"          varchar(512),
    "deletion_time" int8,
    "create_time"   timestamp,
    "update_time"   timestamp,
    "del_flag"      int2,
    CONSTRAINT "_copy_16" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_username" ON "user_0" USING btree (
                                                                     "username" ASC
    );
COMMENT ON COLUMN "user_0"."id" IS 'ID';
COMMENT ON COLUMN "user_0"."username" IS '用户名';
COMMENT ON COLUMN "user_0"."password" IS '密码';
COMMENT ON COLUMN "user_0"."real_name" IS '真实姓名';
COMMENT ON COLUMN "user_0"."phone" IS '手机号';
COMMENT ON COLUMN "user_0"."mail" IS '邮箱';
COMMENT ON COLUMN "user_0"."deletion_time" IS '注销时间戳';
COMMENT ON COLUMN "user_0"."create_time" IS '创建时间';
COMMENT ON COLUMN "user_0"."update_time" IS '修改时间';
COMMENT ON COLUMN "user_0"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "user_1"
(
    "id"            int8 NOT NULL,
    "username"      varchar(256),
    "password"      varchar(512),
    "real_name"     varchar(256),
    "phone"         varchar(128),
    "mail"          varchar(512),
    "deletion_time" int8,
    "create_time"   timestamp,
    "update_time"   timestamp,
    "del_flag"      int2,
    CONSTRAINT "_copy_15" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_username_copy_15" ON "user_1" USING btree (
                                                                             "username" ASC
    );
COMMENT ON COLUMN "user_1"."id" IS 'ID';
COMMENT ON COLUMN "user_1"."username" IS '用户名';
COMMENT ON COLUMN "user_1"."password" IS '密码';
COMMENT ON COLUMN "user_1"."real_name" IS '真实姓名';
COMMENT ON COLUMN "user_1"."phone" IS '手机号';
COMMENT ON COLUMN "user_1"."mail" IS '邮箱';
COMMENT ON COLUMN "user_1"."deletion_time" IS '注销时间戳';
COMMENT ON COLUMN "user_1"."create_time" IS '创建时间';
COMMENT ON COLUMN "user_1"."update_time" IS '修改时间';
COMMENT ON COLUMN "user_1"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "user_10"
(
    "id"            int8 NOT NULL,
    "username"      varchar(256),
    "password"      varchar(512),
    "real_name"     varchar(256),
    "phone"         varchar(128),
    "mail"          varchar(512),
    "deletion_time" int8,
    "create_time"   timestamp,
    "update_time"   timestamp,
    "del_flag"      int2,
    CONSTRAINT "_copy_14" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_username_copy_14" ON "user_10" USING btree (
                                                                              "username" ASC
    );
COMMENT ON COLUMN "user_10"."id" IS 'ID';
COMMENT ON COLUMN "user_10"."username" IS '用户名';
COMMENT ON COLUMN "user_10"."password" IS '密码';
COMMENT ON COLUMN "user_10"."real_name" IS '真实姓名';
COMMENT ON COLUMN "user_10"."phone" IS '手机号';
COMMENT ON COLUMN "user_10"."mail" IS '邮箱';
COMMENT ON COLUMN "user_10"."deletion_time" IS '注销时间戳';
COMMENT ON COLUMN "user_10"."create_time" IS '创建时间';
COMMENT ON COLUMN "user_10"."update_time" IS '修改时间';
COMMENT ON COLUMN "user_10"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "user_11"
(
    "id"            int8 NOT NULL,
    "username"      varchar(256),
    "password"      varchar(512),
    "real_name"     varchar(256),
    "phone"         varchar(128),
    "mail"          varchar(512),
    "deletion_time" int8,
    "create_time"   timestamp,
    "update_time"   timestamp,
    "del_flag"      int2,
    CONSTRAINT "_copy_13" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_username_copy_13" ON "user_11" USING btree (
                                                                              "username" ASC
    );
COMMENT ON COLUMN "user_11"."id" IS 'ID';
COMMENT ON COLUMN "user_11"."username" IS '用户名';
COMMENT ON COLUMN "user_11"."password" IS '密码';
COMMENT ON COLUMN "user_11"."real_name" IS '真实姓名';
COMMENT ON COLUMN "user_11"."phone" IS '手机号';
COMMENT ON COLUMN "user_11"."mail" IS '邮箱';
COMMENT ON COLUMN "user_11"."deletion_time" IS '注销时间戳';
COMMENT ON COLUMN "user_11"."create_time" IS '创建时间';
COMMENT ON COLUMN "user_11"."update_time" IS '修改时间';
COMMENT ON COLUMN "user_11"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "user_12"
(
    "id"            int8 NOT NULL,
    "username"      varchar(256),
    "password"      varchar(512),
    "real_name"     varchar(256),
    "phone"         varchar(128),
    "mail"          varchar(512),
    "deletion_time" int8,
    "create_time"   timestamp,
    "update_time"   timestamp,
    "del_flag"      int2,
    CONSTRAINT "_copy_12" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_username_copy_12" ON "user_12" USING btree (
                                                                              "username" ASC
    );
COMMENT ON COLUMN "user_12"."id" IS 'ID';
COMMENT ON COLUMN "user_12"."username" IS '用户名';
COMMENT ON COLUMN "user_12"."password" IS '密码';
COMMENT ON COLUMN "user_12"."real_name" IS '真实姓名';
COMMENT ON COLUMN "user_12"."phone" IS '手机号';
COMMENT ON COLUMN "user_12"."mail" IS '邮箱';
COMMENT ON COLUMN "user_12"."deletion_time" IS '注销时间戳';
COMMENT ON COLUMN "user_12"."create_time" IS '创建时间';
COMMENT ON COLUMN "user_12"."update_time" IS '修改时间';
COMMENT ON COLUMN "user_12"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "user_13"
(
    "id"            int8 NOT NULL,
    "username"      varchar(256),
    "password"      varchar(512),
    "real_name"     varchar(256),
    "phone"         varchar(128),
    "mail"          varchar(512),
    "deletion_time" int8,
    "create_time"   timestamp,
    "update_time"   timestamp,
    "del_flag"      int2,
    CONSTRAINT "_copy_11" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_username_copy_11" ON "user_13" USING btree (
                                                                              "username" ASC
    );
COMMENT ON COLUMN "user_13"."id" IS 'ID';
COMMENT ON COLUMN "user_13"."username" IS '用户名';
COMMENT ON COLUMN "user_13"."password" IS '密码';
COMMENT ON COLUMN "user_13"."real_name" IS '真实姓名';
COMMENT ON COLUMN "user_13"."phone" IS '手机号';
COMMENT ON COLUMN "user_13"."mail" IS '邮箱';
COMMENT ON COLUMN "user_13"."deletion_time" IS '注销时间戳';
COMMENT ON COLUMN "user_13"."create_time" IS '创建时间';
COMMENT ON COLUMN "user_13"."update_time" IS '修改时间';
COMMENT ON COLUMN "user_13"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "user_14"
(
    "id"            int8 NOT NULL,
    "username"      varchar(256),
    "password"      varchar(512),
    "real_name"     varchar(256),
    "phone"         varchar(128),
    "mail"          varchar(512),
    "deletion_time" int8,
    "create_time"   timestamp,
    "update_time"   timestamp,
    "del_flag"      int2,
    CONSTRAINT "_copy_10" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_username_copy_10" ON "user_14" USING btree (
                                                                              "username" ASC
    );
COMMENT ON COLUMN "user_14"."id" IS 'ID';
COMMENT ON COLUMN "user_14"."username" IS '用户名';
COMMENT ON COLUMN "user_14"."password" IS '密码';
COMMENT ON COLUMN "user_14"."real_name" IS '真实姓名';
COMMENT ON COLUMN "user_14"."phone" IS '手机号';
COMMENT ON COLUMN "user_14"."mail" IS '邮箱';
COMMENT ON COLUMN "user_14"."deletion_time" IS '注销时间戳';
COMMENT ON COLUMN "user_14"."create_time" IS '创建时间';
COMMENT ON COLUMN "user_14"."update_time" IS '修改时间';
COMMENT ON COLUMN "user_14"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "user_15"
(
    "id"            int8 NOT NULL,
    "username"      varchar(256),
    "password"      varchar(512),
    "real_name"     varchar(256),
    "phone"         varchar(128),
    "mail"          varchar(512),
    "deletion_time" int8,
    "create_time"   timestamp,
    "update_time"   timestamp,
    "del_flag"      int2,
    CONSTRAINT "_copy_9" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_username_copy_9" ON "user_15" USING btree (
                                                                             "username" ASC
    );
COMMENT ON COLUMN "user_15"."id" IS 'ID';
COMMENT ON COLUMN "user_15"."username" IS '用户名';
COMMENT ON COLUMN "user_15"."password" IS '密码';
COMMENT ON COLUMN "user_15"."real_name" IS '真实姓名';
COMMENT ON COLUMN "user_15"."phone" IS '手机号';
COMMENT ON COLUMN "user_15"."mail" IS '邮箱';
COMMENT ON COLUMN "user_15"."deletion_time" IS '注销时间戳';
COMMENT ON COLUMN "user_15"."create_time" IS '创建时间';
COMMENT ON COLUMN "user_15"."update_time" IS '修改时间';
COMMENT ON COLUMN "user_15"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "user_2"
(
    "id"            int8 NOT NULL,
    "username"      varchar(256),
    "password"      varchar(512),
    "real_name"     varchar(256),
    "phone"         varchar(128),
    "mail"          varchar(512),
    "deletion_time" int8,
    "create_time"   timestamp,
    "update_time"   timestamp,
    "del_flag"      int2,
    CONSTRAINT "_copy_8" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_username_copy_8" ON "user_2" USING btree (
                                                                            "username" ASC
    );
COMMENT ON COLUMN "user_2"."id" IS 'ID';
COMMENT ON COLUMN "user_2"."username" IS '用户名';
COMMENT ON COLUMN "user_2"."password" IS '密码';
COMMENT ON COLUMN "user_2"."real_name" IS '真实姓名';
COMMENT ON COLUMN "user_2"."phone" IS '手机号';
COMMENT ON COLUMN "user_2"."mail" IS '邮箱';
COMMENT ON COLUMN "user_2"."deletion_time" IS '注销时间戳';
COMMENT ON COLUMN "user_2"."create_time" IS '创建时间';
COMMENT ON COLUMN "user_2"."update_time" IS '修改时间';
COMMENT ON COLUMN "user_2"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "user_3"
(
    "id"            int8 NOT NULL,
    "username"      varchar(256),
    "password"      varchar(512),
    "real_name"     varchar(256),
    "phone"         varchar(128),
    "mail"          varchar(512),
    "deletion_time" int8,
    "create_time"   timestamp,
    "update_time"   timestamp,
    "del_flag"      int2,
    CONSTRAINT "_copy_7" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_username_copy_7" ON "user_3" USING btree (
                                                                            "username" ASC
    );
COMMENT ON COLUMN "user_3"."id" IS 'ID';
COMMENT ON COLUMN "user_3"."username" IS '用户名';
COMMENT ON COLUMN "user_3"."password" IS '密码';
COMMENT ON COLUMN "user_3"."real_name" IS '真实姓名';
COMMENT ON COLUMN "user_3"."phone" IS '手机号';
COMMENT ON COLUMN "user_3"."mail" IS '邮箱';
COMMENT ON COLUMN "user_3"."deletion_time" IS '注销时间戳';
COMMENT ON COLUMN "user_3"."create_time" IS '创建时间';
COMMENT ON COLUMN "user_3"."update_time" IS '修改时间';
COMMENT ON COLUMN "user_3"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "user_4"
(
    "id"            int8 NOT NULL,
    "username"      varchar(256),
    "password"      varchar(512),
    "real_name"     varchar(256),
    "phone"         varchar(128),
    "mail"          varchar(512),
    "deletion_time" int8,
    "create_time"   timestamp,
    "update_time"   timestamp,
    "del_flag"      int2,
    CONSTRAINT "_copy_6" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_username_copy_6" ON "user_4" USING btree (
                                                                            "username" ASC
    );
COMMENT ON COLUMN "user_4"."id" IS 'ID';
COMMENT ON COLUMN "user_4"."username" IS '用户名';
COMMENT ON COLUMN "user_4"."password" IS '密码';
COMMENT ON COLUMN "user_4"."real_name" IS '真实姓名';
COMMENT ON COLUMN "user_4"."phone" IS '手机号';
COMMENT ON COLUMN "user_4"."mail" IS '邮箱';
COMMENT ON COLUMN "user_4"."deletion_time" IS '注销时间戳';
COMMENT ON COLUMN "user_4"."create_time" IS '创建时间';
COMMENT ON COLUMN "user_4"."update_time" IS '修改时间';
COMMENT ON COLUMN "user_4"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "user_5"
(
    "id"            int8 NOT NULL,
    "username"      varchar(256),
    "password"      varchar(512),
    "real_name"     varchar(256),
    "phone"         varchar(128),
    "mail"          varchar(512),
    "deletion_time" int8,
    "create_time"   timestamp,
    "update_time"   timestamp,
    "del_flag"      int2,
    CONSTRAINT "_copy_5" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_username_copy_5" ON "user_5" USING btree (
                                                                            "username" ASC
    );
COMMENT ON COLUMN "user_5"."id" IS 'ID';
COMMENT ON COLUMN "user_5"."username" IS '用户名';
COMMENT ON COLUMN "user_5"."password" IS '密码';
COMMENT ON COLUMN "user_5"."real_name" IS '真实姓名';
COMMENT ON COLUMN "user_5"."phone" IS '手机号';
COMMENT ON COLUMN "user_5"."mail" IS '邮箱';
COMMENT ON COLUMN "user_5"."deletion_time" IS '注销时间戳';
COMMENT ON COLUMN "user_5"."create_time" IS '创建时间';
COMMENT ON COLUMN "user_5"."update_time" IS '修改时间';
COMMENT ON COLUMN "user_5"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "user_6"
(
    "id"            int8 NOT NULL,
    "username"      varchar(256),
    "password"      varchar(512),
    "real_name"     varchar(256),
    "phone"         varchar(128),
    "mail"          varchar(512),
    "deletion_time" int8,
    "create_time"   timestamp,
    "update_time"   timestamp,
    "del_flag"      int2,
    CONSTRAINT "_copy_4" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_username_copy_4" ON "user_6" USING btree (
                                                                            "username" ASC
    );
COMMENT ON COLUMN "user_6"."id" IS 'ID';
COMMENT ON COLUMN "user_6"."username" IS '用户名';
COMMENT ON COLUMN "user_6"."password" IS '密码';
COMMENT ON COLUMN "user_6"."real_name" IS '真实姓名';
COMMENT ON COLUMN "user_6"."phone" IS '手机号';
COMMENT ON COLUMN "user_6"."mail" IS '邮箱';
COMMENT ON COLUMN "user_6"."deletion_time" IS '注销时间戳';
COMMENT ON COLUMN "user_6"."create_time" IS '创建时间';
COMMENT ON COLUMN "user_6"."update_time" IS '修改时间';
COMMENT ON COLUMN "user_6"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "user_7"
(
    "id"            int8 NOT NULL,
    "username"      varchar(256),
    "password"      varchar(512),
    "real_name"     varchar(256),
    "phone"         varchar(128),
    "mail"          varchar(512),
    "deletion_time" int8,
    "create_time"   timestamp,
    "update_time"   timestamp,
    "del_flag"      int2,
    CONSTRAINT "_copy_3" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_username_copy_3" ON "user_7" USING btree (
                                                                            "username" ASC
    );
COMMENT ON COLUMN "user_7"."id" IS 'ID';
COMMENT ON COLUMN "user_7"."username" IS '用户名';
COMMENT ON COLUMN "user_7"."password" IS '密码';
COMMENT ON COLUMN "user_7"."real_name" IS '真实姓名';
COMMENT ON COLUMN "user_7"."phone" IS '手机号';
COMMENT ON COLUMN "user_7"."mail" IS '邮箱';
COMMENT ON COLUMN "user_7"."deletion_time" IS '注销时间戳';
COMMENT ON COLUMN "user_7"."create_time" IS '创建时间';
COMMENT ON COLUMN "user_7"."update_time" IS '修改时间';
COMMENT ON COLUMN "user_7"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "user_8"
(
    "id"            int8 NOT NULL,
    "username"      varchar(256),
    "password"      varchar(512),
    "real_name"     varchar(256),
    "phone"         varchar(128),
    "mail"          varchar(512),
    "deletion_time" int8,
    "create_time"   timestamp,
    "update_time"   timestamp,
    "del_flag"      int2,
    CONSTRAINT "_copy_2" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_username_copy_2" ON "user_8" USING btree (
                                                                            "username" ASC
    );
COMMENT ON COLUMN "user_8"."id" IS 'ID';
COMMENT ON COLUMN "user_8"."username" IS '用户名';
COMMENT ON COLUMN "user_8"."password" IS '密码';
COMMENT ON COLUMN "user_8"."real_name" IS '真实姓名';
COMMENT ON COLUMN "user_8"."phone" IS '手机号';
COMMENT ON COLUMN "user_8"."mail" IS '邮箱';
COMMENT ON COLUMN "user_8"."deletion_time" IS '注销时间戳';
COMMENT ON COLUMN "user_8"."create_time" IS '创建时间';
COMMENT ON COLUMN "user_8"."update_time" IS '修改时间';
COMMENT ON COLUMN "user_8"."del_flag" IS '删除标识 0：未删除 1：已删除';

CREATE TABLE "user_9"
(
    "id"            int8 NOT NULL,
    "username"      varchar(256),
    "password"      varchar(512),
    "real_name"     varchar(256),
    "phone"         varchar(128),
    "mail"          varchar(512),
    "deletion_time" int8,
    "create_time"   timestamp,
    "update_time"   timestamp,
    "del_flag"      int2,
    CONSTRAINT "_copy_1" PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_username_copy_1" ON "user_9" USING btree (
                                                                            "username" ASC
    );
COMMENT ON COLUMN "user_9"."id" IS 'ID';
COMMENT ON COLUMN "user_9"."username" IS '用户名';
COMMENT ON COLUMN "user_9"."password" IS '密码';
COMMENT ON COLUMN "user_9"."real_name" IS '真实姓名';
COMMENT ON COLUMN "user_9"."phone" IS '手机号';
COMMENT ON COLUMN "user_9"."mail" IS '邮箱';
COMMENT ON COLUMN "user_9"."deletion_time" IS '注销时间戳';
COMMENT ON COLUMN "user_9"."create_time" IS '创建时间';
COMMENT ON COLUMN "user_9"."update_time" IS '修改时间';
COMMENT ON COLUMN "user_9"."del_flag" IS '删除标识 0：未删除 1：已删除';

