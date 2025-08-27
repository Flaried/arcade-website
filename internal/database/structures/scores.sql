CREATE TABLE scores (
    id SERIAL PRIMARY KEY,              
    user_id INT NOT NULL,               
    game_id INT NOT NULL,               
    score INT NOT NULL,                 
    achieved_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    is_validated BOOLEAN DEFAULT FALSE,
    -- If a user quits, their scores disappear too
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (game_id) REFERENCES games(id) ON DELETE CASCADE
);
