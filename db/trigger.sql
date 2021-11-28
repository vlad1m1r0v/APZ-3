CREATE OR REPLACE FUNCTION trigger_function()
    RETURNS TRIGGER
LANGUAGE PLPGSQL
AS
$$
DECLARE
_capacity INTEGER;
BEGIN
   -- trigger logic
        _capacity = NEW.capacity;
        UPDATE machines
        SET totaldiskspace = totaldiskspace + _capacity
        WHERE id = NEW.machines_id;
        if OLD.machines_id is not NULL THEN
            UPDATE machines
            SET totaldiskspace = totaldiskspace - _capacity
            WHERE id = OLD.machines_id;
        END IF;
        RETURN NEW;
END;
$$;

CREATE OR REPLACE TRIGGER trigger
    AFTER UPDATE
    ON disks
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_function();