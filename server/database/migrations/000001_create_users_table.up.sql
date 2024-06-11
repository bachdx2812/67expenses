-- Name: users; Type: TABLE; Schema: home_expenses; Owner: -
CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  family_id INT NOT NULL,
  name VARCHAR(255),
  phone VARCHAR(255) NOT NULL,
  encrypted_password VARCHAR(255) DEFAULT NULL,
  created_at TIMESTAMP(6) NOT NULL,
  updated_at TIMESTAMP(6) NOT NULL,
  UNIQUE (phone)
);

CREATE INDEX idx_users_phone ON users (phone);
CREATE INDEX idx_users_family_id ON users (family_id);