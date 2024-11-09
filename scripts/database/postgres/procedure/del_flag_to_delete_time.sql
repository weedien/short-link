CREATE OR REPLACE PROCEDURE update_schema_tables(p_schema_name TEXT)
    LANGUAGE plpgsql
AS
$$
DECLARE
    r               RECORD;
    v_table_name    TEXT;
    has_del_flag    BOOLEAN;
    has_delete_time BOOLEAN;
BEGIN
    -- 循环遍历指定模式下的所有表
    FOR r IN
        SELECT table_name
        FROM information_schema.tables
        WHERE table_schema = p_schema_name
          AND table_type = 'BASE TABLE'
        LOOP
            v_table_name := r.table_name;

            -- 检查是否存在del_flag字段
            SELECT EXISTS (SELECT 1
                           FROM information_schema.columns
                           WHERE table_schema = p_schema_name
                             AND table_name = v_table_name
                             AND column_name = 'del_flag')
            INTO has_del_flag;

            -- 检查是否存在delete_time字段
            SELECT EXISTS (SELECT 1
                           FROM information_schema.columns
                           WHERE table_schema = p_schema_name
                             AND table_name = v_table_name
                             AND column_name = 'delete_time')
            INTO has_delete_time;

            -- 如果存在del_flag字段，则删除它
            IF has_del_flag THEN
                EXECUTE format('ALTER TABLE %I.%I DROP COLUMN del_flag', p_schema_name, v_table_name);
            END IF;

            -- 如果不存在delete_time字段，则添加它并加上注释
            IF NOT has_delete_time THEN
                EXECUTE format('ALTER TABLE %I.%I ADD COLUMN delete_time TIMESTAMP', p_schema_name, v_table_name);
                EXECUTE format('COMMENT ON COLUMN %I.%I.delete_time IS ''删除时间''', p_schema_name, v_table_name);
            END IF;

            RAISE NOTICE 'Updated table %', v_table_name;
        END LOOP;
END
$$;


-- 调用存储过程
CALL update_schema_tables('link_single');