CREATE OR REPLACE FUNCTION add_classgroup_to_current_schoolyear()
RETURNS TRIGGER AS $$
DECLARE
    current_schoolyear_id INTEGER;
BEGIN
    -- Find the ID of the current SchoolYear
    SELECT id INTO current_schoolyear_id FROM school_years WHERE is_current = TRUE LIMIT 1;

    -- If a current SchoolYear exists, create a relation in the join table
    IF FOUND THEN
        INSERT INTO school_year_class_group (school_year_id, class_group_id) VALUES (current_schoolyear_id, NEW.id);
    END IF;

    -- Return the new row
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER add_classgroup_after_insert
AFTER INSERT ON class_groups
FOR EACH ROW
EXECUTE FUNCTION add_classgroup_to_current_schoolyear();
