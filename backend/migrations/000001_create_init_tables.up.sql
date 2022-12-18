
CREATE TABLE articles (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    created_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    title VARCHAR (255) NOT NULL,
    content TEXT NOT NULL,
    status VARCHAR (150) NOT NULL,
    category VARCHAR (150) NOT NULL
);

CREATE TABLE logging (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    user_id INT NOT NULL,
    slug VARCHAR (255) NOT NULL,
    action_status INT NOT NULL,
    action_attrs JSON 
);

-- Add indexes
CREATE INDEX active_article ON article (id) ;
CREATE INDEX active_logging ON logging (id) ;

