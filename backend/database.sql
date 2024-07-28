CREATE TYPE content_type AS ENUM ('ppt', 'video');

CREATE TABLE content (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  content_type content_type NOT NULL,
  content_url VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  expires_at TIMESTAMP
);

