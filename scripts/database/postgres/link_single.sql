CREATE OR REPLACE FUNCTION cs_timestamp() RETURNS TRIGGER AS
$$
BEGIN
    NEW.update_time = CURRENT_TIMESTAMP;
    RETURN NEW;
END
$$ LANGUAGE plpgsql;

DROP TABLE IF EXISTS "group";
CREATE TABLE "group"
(
    "id"          serial         NOT NULL,
    "gid"         VARCHAR(32)  NOT NULL UNIQUE CHECK (gid <> ''),
    "name"        VARCHAR(64)  NOT NULL,
    "username"    VARCHAR(256) NOT NULL UNIQUE CHECK (username <> ''),
    "sort_order"  INT4,
    "create_time" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "update_time" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "del_flag"    BOOLEAN   DEFAULT FALSE,
    PRIMARY KEY ("id")
);

CREATE TRIGGER update_table_timestamp
    BEFORE INSERT
        OR UPDATE
    ON "group"
    FOR EACH ROW
EXECUTE FUNCTION cs_timestamp();

CREATE INDEX "idx_username" ON "group" USING btree ("username" ASC);
COMMENT ON COLUMN "group"."id" IS 'ID';
COMMENT ON COLUMN "group"."gid" IS '分组标识';
COMMENT ON COLUMN "group"."name" IS '分组名称';
COMMENT ON COLUMN "group"."username" IS '创建分组用户名';
COMMENT ON COLUMN "group"."sort_order" IS '分组排序';
COMMENT ON COLUMN "group"."create_time" IS '创建时间';
COMMENT ON COLUMN "group"."update_time" IS '修改时间';
COMMENT ON COLUMN "group"."del_flag" IS '删除标识 0：未删除 1：已删除';

DROP TABLE IF EXISTS "group_unique";
CREATE TABLE "group_unique"
(
    "id"  INT8        NOT NULL,
    "gid" VARCHAR(32) NOT NULL UNIQUE CHECK (gid <> ''),
    PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "idx_unique_gid" ON "group_unique" USING btree ("gid" ASC);
COMMENT ON COLUMN "group_unique"."id" IS 'ID';
COMMENT ON COLUMN "group_unique"."gid" IS '分组标识';

DROP TABLE IF EXISTS "user";
CREATE TABLE "user"
(
    "id"            INT8         NOT NULL,
    "username"      VARCHAR(256) NOT NULL UNIQUE CHECK (username <> ''),
    "password"      VARCHAR(512),
    "real_name"     VARCHAR(64),
    "phone"         VARCHAR(128),
    "mail"          VARCHAR(512),
    "deletion_time" TIMESTAMP,
    "create_time"   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "update_time"   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "del_flag"      BOOLEAN   DEFAULT FALSE,
    PRIMARY KEY ("id")
);

CREATE TRIGGER update_table_timestamp
    BEFORE INSERT
        OR UPDATE
    ON "user"
    FOR EACH ROW
EXECUTE FUNCTION cs_timestamp();

CREATE UNIQUE INDEX "idx_unique_username" ON "user" USING btree ("username" ASC);
COMMENT ON COLUMN "user"."id" IS 'ID';
COMMENT ON COLUMN "user"."username" IS '用户名';
COMMENT ON COLUMN "user"."password" IS '密码';
COMMENT ON COLUMN "user"."real_name" IS '真实姓名';
COMMENT ON COLUMN "user"."phone" IS '手机号';
COMMENT ON COLUMN "user"."mail" IS '邮箱';
COMMENT ON COLUMN "user"."deletion_time" IS '注销时间戳';
COMMENT ON COLUMN "user"."create_time" IS '创建时间';
COMMENT ON COLUMN "user"."update_time" IS '修改时间';
COMMENT ON COLUMN "user"."del_flag" IS '删除标识 0：未删除 1：已删除';

DROP TABLE IF EXISTS "link";
CREATE TABLE "link"
(
    "id"              INT8 NOT NULL,
    "domain"          VARCHAR(128),
    "short_uri"       VARCHAR(8),
    "full_short_url"  VARCHAR(128),
    "origin_url"      VARCHAR(1024),
    "click_num"       INT4,
    "gid"             VARCHAR(32),
    "favicon"         VARCHAR(256),
    "enable_status"   BOOLEAN,
    "created_type"    INT2,
    "valid_date_type" INT2,
    "valid_date"      TIMESTAMP,
    "describe"        VARCHAR(1024),
    "total_pv"        INT4,
    "total_uv"        INT4,
    "total_uip"       INT4,
    "create_time"     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "update_time"     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "del_time"        TIMESTAMP,
    "del_flag"        BOOLEAN   DEFAULT FALSE,
    PRIMARY KEY ("id")
);

CREATE TRIGGER update_table_timestamp
    BEFORE INSERT
        OR UPDATE
    ON "link"
    FOR EACH ROW
EXECUTE FUNCTION cs_timestamp();

CREATE UNIQUE INDEX "idx_unique_full-short-url" ON "link" USING btree ("full_short_url" ASC, "del_time" ASC);
COMMENT ON COLUMN "link"."id" IS 'ID';
COMMENT ON COLUMN "link"."domain" IS '域名';
COMMENT ON COLUMN "link"."short_uri" IS '短链接';
COMMENT ON COLUMN "link"."full_short_url" IS '完整短链接';
COMMENT ON COLUMN "link"."origin_url" IS '原始链接';
COMMENT ON COLUMN "link"."click_num" IS '点击量';
COMMENT ON COLUMN "link"."gid" IS '分组标识';
COMMENT ON COLUMN "link"."favicon" IS '网站图标';
COMMENT ON COLUMN "link"."enable_status" IS '启用标识';
COMMENT ON COLUMN "link"."created_type" IS '创建类型 0：接口创建 1：控制台创建';
COMMENT ON COLUMN "link"."valid_date_type" IS '有效期类型 0：永久有效 1：自定义';
COMMENT ON COLUMN "link"."valid_date" IS '有效期';
COMMENT ON COLUMN "link"."describe" IS '描述';
COMMENT ON COLUMN "link"."total_pv" IS '历史PV';
COMMENT ON COLUMN "link"."total_uv" IS '历史UV';
COMMENT ON COLUMN "link"."total_uip" IS '历史UIP';
COMMENT ON COLUMN "link"."create_time" IS '创建时间';
COMMENT ON COLUMN "link"."update_time" IS '修改时间';
COMMENT ON COLUMN "link"."del_time" IS '删除时间戳';
COMMENT ON COLUMN "link"."del_flag" IS '删除标识 0：未删除 1：已删除';

DROP TABLE IF EXISTS "link_access_logs";
CREATE TABLE "link_access_logs"
(
    "id"             INT8 NOT NULL,
    "full_short_url" VARCHAR(128),
    "user"           VARCHAR(64),
    "ip"             VARCHAR(64),
    "browser"        VARCHAR(64),
    "os"             VARCHAR(64),
    "network"        VARCHAR(64),
    "device"         VARCHAR(64),
    "locale"         VARCHAR(256),
    "create_time"    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "update_time"    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "del_flag"       BOOLEAN   DEFAULT FALSE,
    PRIMARY KEY ("id")
);

CREATE TRIGGER update_table_timestamp
    BEFORE INSERT
        OR UPDATE
    ON "link_access_logs"
    FOR EACH ROW
EXECUTE FUNCTION cs_timestamp();

CREATE INDEX "idx_full_short_url" ON "link_access_logs" USING btree ("full_short_url" ASC);
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

DROP TABLE IF EXISTS "link_access_stats";
CREATE TABLE "link_access_stats"
(
    "id"             INT8 NOT NULL,
    "full_short_url" VARCHAR(128),
    "date"           DATE,
    "pv"             INT4,
    "uv"             INT4,
    "uip"            INT4,
    "hour"           INT4,
    "weekday"        INT4,
    "create_time"    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "update_time"    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "del_flag"       BOOLEAN   DEFAULT FALSE,
    PRIMARY KEY ("id")
);

CREATE TRIGGER update_table_timestamp
    BEFORE INSERT
        OR UPDATE
    ON "link_access_stats"
    FOR EACH ROW
EXECUTE FUNCTION cs_timestamp();

CREATE UNIQUE INDEX "idx_unique_access_stats" ON "link_access_stats" USING btree ("full_short_url" ASC, "date" ASC, "hour" ASC);
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

DROP TABLE IF EXISTS "link_browser_stats";
CREATE TABLE "link_browser_stats"
(
    "id"             INT8 NOT NULL,
    "full_short_url" VARCHAR(128),
    "date"           DATE,
    "cnt"            INT4,
    "browser"        VARCHAR(64),
    "create_time"    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "update_time"    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "del_flag"       BOOLEAN   DEFAULT FALSE,
    PRIMARY KEY ("id")
);

CREATE TRIGGER update_table_timestamp
    BEFORE INSERT
        OR UPDATE
    ON "link_browser_stats"
    FOR EACH ROW
EXECUTE FUNCTION cs_timestamp();

CREATE UNIQUE INDEX "idx_unique_browser_stats" ON "link_browser_stats" USING btree ("full_short_url" ASC, "date" ASC, "browser" ASC);
COMMENT ON COLUMN "link_browser_stats"."id" IS 'ID';
COMMENT ON COLUMN "link_browser_stats"."full_short_url" IS '完整短链接';
COMMENT ON COLUMN "link_browser_stats"."date" IS '日期';
COMMENT ON COLUMN "link_browser_stats"."cnt" IS '访问量';
COMMENT ON COLUMN "link_browser_stats"."browser" IS '浏览器';
COMMENT ON COLUMN "link_browser_stats"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_browser_stats"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_browser_stats"."del_flag" IS '删除标识 0：未删除 1：已删除';

DROP TABLE IF EXISTS "link_device_stats";
CREATE TABLE "link_device_stats"
(
    "id"             INT8 NOT NULL,
    "full_short_url" VARCHAR(128),
    "date"           DATE,
    "cnt"            INT4,
    "device"         VARCHAR(64),
    "create_time"    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "update_time"    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "del_flag"       BOOLEAN   DEFAULT FALSE,
    PRIMARY KEY ("id")
);

CREATE TRIGGER update_table_timestamp
    BEFORE INSERT
        OR UPDATE
    ON "link_device_stats"
    FOR EACH ROW
EXECUTE FUNCTION cs_timestamp();

CREATE UNIQUE INDEX "idx_unique_device_stats" ON "link_device_stats" USING btree ("full_short_url" ASC, "date" ASC, "device" ASC);
COMMENT ON COLUMN "link_device_stats"."id" IS 'ID';
COMMENT ON COLUMN "link_device_stats"."full_short_url" IS '完整短链接';
COMMENT ON COLUMN "link_device_stats"."date" IS '日期';
COMMENT ON COLUMN "link_device_stats"."cnt" IS '访问量';
COMMENT ON COLUMN "link_device_stats"."device" IS '设备';
COMMENT ON COLUMN "link_device_stats"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_device_stats"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_device_stats"."del_flag" IS '删除标识 0：未删除 1：已删除';

DROP TABLE IF EXISTS "link_goto";
CREATE TABLE "link_goto"
(
    "id"             INT8 NOT NULL,
    "gid"            VARCHAR(32),
    "full_short_url" VARCHAR(128),
    PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "idx_unique_goto" ON "link_goto" USING btree ("gid" ASC, "full_short_url" ASC);
COMMENT ON COLUMN "link_goto"."id" IS 'ID';
COMMENT ON COLUMN "link_goto"."gid" IS '分组标识';
COMMENT ON COLUMN "link_goto"."full_short_url" IS '完整短链接';

DROP TABLE IF EXISTS "link_locale_stats";
CREATE TABLE "link_locale_stats"
(
    "id"             INT8 NOT NULL,
    "full_short_url" VARCHAR(128),
    "date"           DATE,
    "cnt"            INT4,
    "province"       VARCHAR(64),
    "city"           VARCHAR(64),
    "adcode"         VARCHAR(64),
    "country"        VARCHAR(64),
    "create_time"    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "update_time"    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "del_flag"       BOOLEAN   DEFAULT FALSE,
    PRIMARY KEY ("id")
);

CREATE TRIGGER update_table_timestamp
    BEFORE INSERT
        OR UPDATE
    ON "link_locale_stats"
    FOR EACH ROW
EXECUTE FUNCTION cs_timestamp();

CREATE UNIQUE INDEX "idx_unique_locale_stats" ON "link_locale_stats" USING btree ("full_short_url" ASC, "date" ASC, "adcode" ASC, "province" ASC);
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

DROP TABLE IF EXISTS "link_network_stats";
CREATE TABLE "link_network_stats"
(
    "id"             INT8 NOT NULL,
    "full_short_url" VARCHAR(128),
    "date"           DATE,
    "cnt"            INT4,
    "network"        VARCHAR(64),
    "create_time"    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "update_time"    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "del_flag"       BOOLEAN   DEFAULT FALSE,
    PRIMARY KEY ("id")
);

CREATE TRIGGER update_table_timestamp
    BEFORE INSERT
        OR UPDATE
    ON "link_network_stats"
    FOR EACH ROW
EXECUTE FUNCTION cs_timestamp();

CREATE UNIQUE INDEX "idx_unique_network_stats" ON "link_network_stats" USING btree ("full_short_url" ASC, "date" ASC, "network" ASC);
COMMENT ON COLUMN "link_network_stats"."id" IS 'ID';
COMMENT ON COLUMN "link_network_stats"."full_short_url" IS '完整短链接';
COMMENT ON COLUMN "link_network_stats"."date" IS '日期';
COMMENT ON COLUMN "link_network_stats"."cnt" IS '访问量';
COMMENT ON COLUMN "link_network_stats"."network" IS '网络';
COMMENT ON COLUMN "link_network_stats"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_network_stats"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_network_stats"."del_flag" IS '删除标识 0：未删除 1：已删除';

DROP TABLE IF EXISTS "link_os_stats";
CREATE TABLE "link_os_stats"
(
    "id"             INT8 NOT NULL,
    "full_short_url" VARCHAR(128),
    "date"           DATE,
    "cnt"            INT4,
    "os"             VARCHAR(64),
    "create_time"    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "update_time"    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "del_flag"       BOOLEAN   DEFAULT FALSE,
    PRIMARY KEY ("id")
);

CREATE TRIGGER update_table_timestamp
    BEFORE INSERT
        OR UPDATE
    ON "link_os_stats"
    FOR EACH ROW
EXECUTE FUNCTION cs_timestamp();

CREATE UNIQUE INDEX "idx_unique_os_stats" ON "link_os_stats" USING btree ("full_short_url" ASC, "date" ASC, "os" ASC);
COMMENT ON COLUMN "link_os_stats"."id" IS 'ID';
COMMENT ON COLUMN "link_os_stats"."full_short_url" IS '完整短链接';
COMMENT ON COLUMN "link_os_stats"."date" IS '日期';
COMMENT ON COLUMN "link_os_stats"."cnt" IS '访问量';
COMMENT ON COLUMN "link_os_stats"."os" IS '操作系统';
COMMENT ON COLUMN "link_os_stats"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_os_stats"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_os_stats"."del_flag" IS '删除标识 0：未删除 1：已删除';

DROP TABLE IF EXISTS "link_stats_today";
CREATE TABLE "link_stats_today"
(
    "id"             INT8 NOT NULL,
    "full_short_url" VARCHAR(128),
    "date"           DATE,
    "today_pv"       INT4,
    "today_uv"       INT4,
    "today_uip"      INT4,
    "create_time"    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "update_time"    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "del_flag"       BOOLEAN   DEFAULT FALSE,
    PRIMARY KEY ("id")
);

CREATE TRIGGER update_table_timestamp
    BEFORE INSERT
        OR UPDATE
    ON "link_stats_today"
    FOR EACH ROW
EXECUTE FUNCTION cs_timestamp();

CREATE UNIQUE INDEX "idx_unique_today_stats" ON "link_stats_today" USING btree ("full_short_url" ASC, "date" ASC);
COMMENT ON COLUMN "link_stats_today"."id" IS 'ID';
COMMENT ON COLUMN "link_stats_today"."full_short_url" IS '短链接';
COMMENT ON COLUMN "link_stats_today"."date" IS '日期';
COMMENT ON COLUMN "link_stats_today"."today_pv" IS '今日PV';
COMMENT ON COLUMN "link_stats_today"."today_uv" IS '今日UV';
COMMENT ON COLUMN "link_stats_today"."today_uip" IS '今日IP数';
COMMENT ON COLUMN "link_stats_today"."create_time" IS '创建时间';
COMMENT ON COLUMN "link_stats_today"."update_time" IS '修改时间';
COMMENT ON COLUMN "link_stats_today"."del_flag" IS '删除标识 0：未删除 1：已删除';

-- 默认用户
INSERT INTO "group" ("id", "gid", "name", "username", "sort_order", "create_time", "update_time", "del_flag")
VALUES (1752265619253805057, 'tSUBMP', '默认分组', 'admin', 0, '2024-07-31 21:00:00', '2024-07-31 21:00:00', false);

INSERT INTO "user" ("id", "username", "password", "real_name", "phone", "mail", "deletion_time", "create_time",
                    "update_time", "del_flag")
VALUES (1752265616481370113, 'admin', 'admin123456', 'admin', 'yKZz0xLyjNb9LSCOCfJD4w==', '02/9oF/nWTBK0cM8UPtCOw==',
        NULL, '2024-07-31 21:00:00', '2024-07-31 21:00:00', false);