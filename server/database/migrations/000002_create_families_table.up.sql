-- Name: families; Type: TABLE; Schema: home_expenses; Owner: -
CREATE TABLE families (
  id BIGSERIAL PRIMARY KEY,
  name TEXT,
  created_at TIMESTAMP(6) NOT NULL,
  updated_at TIMESTAMP(6) NOT NULL
);
