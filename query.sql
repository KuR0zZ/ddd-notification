-- Create notifications table

CREATE TABLE Notifications (
	id uuid DEFAULT gen_random_uuid(),
	email VARCHAR NOT NULL,
	message TEXT NOT NULL,
	type VARCHAR NOT NULL,
	is_sent BOOLEAN NOT NULL,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (id)
);

CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_notification_modtime
BEFORE UPDATE ON Notifications
FOR EACH ROW
EXECUTE FUNCTION update_modified_column();