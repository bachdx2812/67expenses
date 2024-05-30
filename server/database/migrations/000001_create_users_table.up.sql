-- Name: users; Type: TABLE; Schema: home_expenses; Owner: -
CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(255),
  phone VARCHAR(255) NOT NULL,
  encrypted_password VARCHAR(255) DEFAULT NULL,
  created_at TIMESTAMP(6) NOT NULL,
  updated_at TIMESTAMP(6) NOT NULL,
  UNIQUE (phone)
);