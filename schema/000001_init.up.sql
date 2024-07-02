CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL CHECK (LENGTH(content) <= 2000),
    user_id INT NOT NULL,
    blocked BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL CHECK (LENGTH(content) <= 2000),
    user_id INT NOT NULL,
    post_id INT NOT NULL,
    parent_id INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    constraint fk_post
        foreign key(post_id)
            references posts(id)
            on delete cascade
);