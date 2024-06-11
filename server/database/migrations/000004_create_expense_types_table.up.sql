-- Name: expenses_types; Type: TABLE; Schema: home_expenses; Owner: -
CREATE TABLE expense_types (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(255),
  created_at TIMESTAMP(6) NOT NULL,
  updated_at TIMESTAMP(6) NOT NULL
);