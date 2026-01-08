CREATE TABLE migration (
  version INT PRIMARY KEY,
  name TEXT NOT NULL,
  query TEXT NOT NULL,
  applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
