-- Create notifications table

CREATE TABLE Notifications (
	id uuid DEFAULT gen_random_uuid(),
	email VARCHAR NOT NULL,
	message TEXT NOT NULL,
	type VARCHAR NOT NULL,
	status VARCHAR NOT NULL,
	PRIMARY KEY (id)
);