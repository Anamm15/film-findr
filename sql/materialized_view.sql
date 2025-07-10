-- Materialzed view for review

CREATE MATERIALIZED VIEW review_mingguan AS
SELECT
  date_trunc('week', created_at) AS minggu,
  COUNT(*) AS jumlah_review
FROM review
GROUP BY 1
ORDER BY 1;

-- indexing for better performance
CREATE INDEX idx_review_created_at ON review(created_at);
CREATE UNIQUE INDEX idx_review_mingguan ON review_mingguan(minggu);

-- refreshing
REFRESH MATERIALIZED VIEW CONCURRENTLY review_mingguan;
-- or 
0 0 * * 1 psql -U postgres -d database_name -c "REFRESH MATERIALIZED VIEW review_mingguan"



