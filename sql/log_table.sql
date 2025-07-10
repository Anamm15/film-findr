--  log watchlist for top film all time
CREATE TABLE log_watchlist (
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users(id),
  film_id INTEGER REFERENCES film(id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION log_watchlist_trigger()
RETURNS trigger AS $$
BEGIN
  INSERT INTO log_watchlist (user_id, film_id, action)
  VALUES (NEW.user_id, NEW.film_id, 'add');
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER trg_log_watchlist
AFTER INSERT ON watchlist
FOR EACH ROW
EXECUTE FUNCTION log_watchlist_trigger();


CREATE MATERIALIZED VIEW top_film_watchlist AS
SELECT
  film_id,
  COUNT(*) AS jumlah_add
FROM log_watchlist
GROUP BY film_id
ORDER BY jumlah_add DESC;


SELECT * FROM top_film_watchlist ORDER BY jumlah_add DESC LIMIT 10;

CREATE UNIQUE INDEX IF NOT EXISTS idx_top_film_watchlist_film_id
ON top_film_watchlist(film_id);

REFRESH MATERIALIZED VIEW CONCURRENTLY top_film_watchlist;



-- trending film
CREATE MATERIALIZED VIEW trending_film_weekly AS
SELECT
  film_id,
  COUNT(*) AS jumlah_add,
  MAX(created_at) AS last_added_at
FROM log_watchlist
WHERE action = 'add'
  AND created_at >= date_trunc('week', NOW())
GROUP BY film_id
ORDER BY jumlah_add DESC;


CREATE UNIQUE INDEX IF NOT EXISTS idx_trending_film_film_id
ON trending_film_weekly(film_id);


REFRESH MATERIALIZED VIEW CONCURRENTLY trending_film_weekly;
