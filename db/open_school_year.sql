CREATE OR REPLACE FUNCTION open_new_school_year(year_name TEXT, start_date TIMESTAMP, end_date TIMESTAMP) RETURNS VOID AS $$
DECLARE
    new_year_id INT;
BEGIN
    -- Check if school year exists
    IF NOT EXISTS (SELECT 1 FROM school_years WHERE name = year_name) THEN
        -- Make all previous years archived
        UPDATE school_years SET is_current = false;
        -- Add new school year
        INSERT INTO school_years ("name", "start", "end", is_current, created_at, updated_at) VALUES (year_name, start_date, end_date, true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING id INTO new_year_id;

        -- Insert records into school_year_class_group for class groups with education_year < 8
        INSERT INTO school_year_class_group (school_year_id, class_group_id)
        SELECT new_year_id, id FROM class_groups WHERE education_year < 8;

        -- Increment education_year by one for each class group
        UPDATE class_groups SET education_year = education_year + 1 WHERE education_year < 8;

    ELSE
        RAISE EXCEPTION 'School year with this name exists';
    END IF;
END;
$$ LANGUAGE plpgsql;
