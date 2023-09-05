CREATE TABLE bills (
  id                     BIGSERIAL PRIMARY KEY,
  customer               INTEGER   NOT NULL,
  time_period_in_seconds INTEGER   NOT NULL,
  closed_at              TIMESTAMP
);

CREATE TABLE bill_charges (
  id      BIGSERIAL PRIMARY KEY,
  bill_id INTEGER   NOT NULL REFERENCES bills(id) ON DELETE CASCADE,
  amount  DECIMAL   NOT NULL
)
