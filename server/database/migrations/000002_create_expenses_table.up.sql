-- Name: expenses; Type: TABLE; Schema: home_expenses; Owner: -
CREATE TABLE expenses (
  id BIGSERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  content TEXT,
  date TIMESTAMP(6) NOT NULL,
  created_at TIMESTAMP(6) NOT NULL,
  updated_at TIMESTAMP(6) NOT NULL
);