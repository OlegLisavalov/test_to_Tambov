CREATE TABLE news (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL
);

CREATE TABLE NewsCategories (
    news_id BIGINT NOT NULL,
    category_id BIGINT NOT NULL,
    PRIMARY KEY (news_id, category_id)
);
