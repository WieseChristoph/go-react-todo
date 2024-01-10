CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE public.user (
  id BIGINT PRIMARY KEY,
  username VARCHAR(255) NOT NULL,
  global_name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  avatar VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON public.user
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE public.todo (
	id SERIAL PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	description TEXT NOT NULL,
	status BOOL NOT NULL DEFAULT true,
  user_id BIGINT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

  FOREIGN KEY (user_id) REFERENCES public.user(id)
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON public.todo
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE public.session (
	id SERIAL PRIMARY KEY,
	token UUID NOT NULL,
	user_id BIGINT NOT NULL,
	expires_at TIMESTAMP NOT NULL,

  FOREIGN KEY (user_id) REFERENCES public.user(id)
);