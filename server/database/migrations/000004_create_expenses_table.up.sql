-- Name: expenses; Type: TABLE; Schema: home_expenses; Owner: -
CREATE TABLE expenses (
  id BIGSERIAL PRIMARY KEY,
  family_id INT NOT NULL,
  user_id INT NOT NULL,
  expense_type_id INT NOT NULL,
  content TEXT,
  date TIMESTAMP(6) NOT NULL,
  created_at TIMESTAMP(6) NOT NULL,
  updated_at TIMESTAMP(6) NOT NULL
);

CREATE INDEX idx_expenses_user_id ON expenses (user_id);
CREATE INDEX idx_expenses_family_id ON expenses (family_id);
CREATE INDEX idx_expenses_expense_type_id ON expenses (expense_type_id);
CREATE INDEX idx_expenses_date ON expenses (date);