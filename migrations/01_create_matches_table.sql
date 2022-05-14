BEGIN;

CREATE TABLE IF NOT EXISTS matches (
      id varchar(64) primary key not null,
      state jsonb
);

COMMIT;