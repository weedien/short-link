CREATE OR REPLACE FUNCTION convert_id_int8_serial(schema_name TEXT) RETURNS VOID AS
$$
DECLARE
    table_record  RECORD;
    column_record RECORD;
    seq_name      TEXT;
BEGIN
    -- Loop through each table in the specified schema
    FOR table_record IN
        SELECT table_name
        FROM information_schema.tables
        WHERE table_schema = schema_name
          AND table_type = 'BASE TABLE'
        LOOP
            -- Loop through each column that is of type int8 and is named id
            FOR column_record IN
                SELECT column_name
                FROM information_schema.columns
                WHERE table_schema = schema_name
                  AND table_name = table_record.table_name
                  AND column_name = 'id'
                  AND data_type = 'bigint'
                  AND column_default IS NULL
                LOOP
                    -- Generate sequence name based on table and column name
                    seq_name := table_record.table_name || '_' || column_record.column_name || '_seq';

                    -- Create a new sequence
                    EXECUTE format('CREATE SEQUENCE %I.%I', schema_name, seq_name);

                    -- Set default originalUrl for the id column to the next originalUrl of the new sequence
                    EXECUTE format('ALTER TABLE %I.%I ALTER COLUMN %I SET DEFAULT nextval(%L)',
                                   schema_name, table_record.table_name, column_record.column_name,
                                   schema_name || '.' || seq_name);

                    -- Set the id column to use the new sequence
                    EXECUTE format('ALTER SEQUENCE %I.%I OWNED BY %I.%I.%I',
                                   schema_name, seq_name, schema_name, table_record.table_name,
                                   column_record.column_name);

                    -- Optionally, set the sequence to start from the max id originalUrl to avoid conflicts
                    EXECUTE format('SELECT setval(%L, COALESCE(MAX(%I), 0) + 1, false) FROM %I.%I',
                                   schema_name || '.' || seq_name, column_record.column_name, schema_name,
                                   table_record.table_name);

                    RAISE NOTICE 'Converted % to serial in table %', column_record.column_name, table_record.table_name;
                END LOOP;
        END LOOP;
END;
$$ LANGUAGE plpgsql;

SELECT convert_id_int8_serial('link_single');

-- 修改 t_link 表的 id 字段类型
create sequence t_link_id_seq;
alter table t_link alter column id set default nextval('t_link_id_seq' :: regclass);
alter sequence t_link_id_seq owned by t_link.id;
select setval('t_link_id_seq', coalesce(max(id), 0) + 1, false) from t_link;


select format('ALTER SEQUENCE %I.%I OWNED BY %I.%I.%I', 'link_single', 't_link_id_seq', 'link_single', 't_link', 'id');