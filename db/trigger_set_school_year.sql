CREATE OR REPLACE FUNCTION set_school_year()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW."school_year_id" IS NULL THEN
        -- Pobierz ID roku szkolnego, który jest oznaczony jako aktualny
        SELECT id INTO NEW."school_year_id"
        FROM school_years
        WHERE is_current = true;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_set_mark_school_year
BEFORE INSERT ON marks
FOR EACH ROW
EXECUTE FUNCTION set_school_year();

CREATE TRIGGER trigger_set_lesson_school_year
BEFORE INSERT ON lessons
FOR EACH ROW
EXECUTE FUNCTION set_school_year();

