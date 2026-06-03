Used EXPLAIN ANALYZE to inspect a therapist productivity query. PostgreSQL performed a sequential scan on the small therapy_visits table, used the therapists primary key index for join lookups, memoized repeated therapist_id lookups, then aggregated results by therapist name.
EXPLAIN ANALYZE
Seq Scan
- Reads every row in a table.
Index Scan
- Uses an index to locate rows.
Nested Loop
- Join strategy that looks up matching rows.
HashAggregate
- Used for GROUP BY operations.
Memoize
- Caches repeated lookups during joins.
