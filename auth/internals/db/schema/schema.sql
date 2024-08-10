CREATE   TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    username TEXT NOT NULL UNIQUE
);


CREATE TABLE IF NOT EXISTS   refresh_tokens (
    token_id SERIAL PRIMARY KEY,
    user_id TEXT NOT NULL,
    token TEXT NOT NULL,
    expires_at INTEGER NOT NULL
);


INSERT INTO users (id, name, email, password, username) VALUES ('1b4e28ba-2fa1-11d2-883f-0016d3cca427', 'João Silva', 'joao.silva@example.com', 'b2e98ad6f6eb8508dd6a14cfa704bad7f05f6fb1c8c7e0fcd494b3a2a23f221a', 'joaosilva') ON CONFLICT (email, username) DO NOTHING;
INSERT INTO users (id, name, email, password, username) VALUES ('1b4e28ba-2fa1-11d2-883f-0016d3cca428', 'Maria Souza', 'maria.souza@example.com', 'b3e98ad6f6eb8508dd6a14cfa704bad7f05f6fb1c8c7e0fcd494b3a2a23f221b', 'mariasouza') ON CONFLICT (email, username) DO NOTHING;
INSERT INTO users (id, name, email, password, username) VALUES ('1b4e28ba-2fa1-11d2-883f-0016d3cca429', 'Pedro Santos', 'pedro.santos@example.com', 'b4e98ad6f6eb8508dd6a14cfa704bad7f05f6fb1c8c7e0fcd494b3a2a23f221c', 'pedrosantos') ON CONFLICT (email, username) DO NOTHING;
INSERT INTO users (id, name, email, password, username) VALUES ('1b4e28ba-2fa1-11d2-883f-0016d3cca42a', 'Ana Oliveira', 'ana.oliveira@example.com', 'b5e98ad6f6eb8508dd6a14cfa704bad7f05f6fb1c8c7e0fcd494b3a2a23f221d', 'anaoliveira') ON CONFLICT (email, username) DO NOTHING;
INSERT INTO users (id, name, email, password, username) VALUES ('1b4e28ba-2fa1-11d2-883f-0016d3cca42b', 'Lucas Costa', 'lucas.costa@example.com', 'b6e98ad6f6eb8508dd6a14cfa704bad7f05f6fb1c8c7e0fcd494b3a2a23f221e', 'lucascosta') ON CONFLICT (email, username) DO NOTHING;
