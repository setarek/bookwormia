CREATE TABLE IF NOT EXISTS book_comments (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    book_id BIGINT NOT NULL,
    comment TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (book_id) REFERENCES books(id),
    CONSTRAINT unique_user_book_comment UNIQUE (user_id, book_id)
);