-- name: CreateFeedFollow :one
WITH cte AS 
(
	INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
	VALUES (
		$1,
		$2,
		$3,
		$4,
		$5
	)
	RETURNING *
) SELECT (
	cte.id,
	cte.created_at,
	cte.updated_at,
	cte.user_id,
	cte.feed_id,
	users.name,
	feeds.name
)
FROM cte
	JOIN users ON cte.user_id = users.id
	JOIN feeds ON cte.feed_id = feeds.id;

-- name: GetFeedFollowsForUser :many
SELECT 
	feed_follows.*, 
 	users.name AS user_name,
	feeds.name AS feed_name
FROM feed_follows
JOIN users ON feed_follows.user_id = users.id
JOIN feeds ON feed_follows.feed_id = feeds.id
WHERE users.id = $1;
