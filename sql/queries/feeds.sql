-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
	$1,
	$2,
	$3,
	$4,
	$5,
	$6
)
RETURNING *;
--

-- name: GetFeeds :many
SELECT 
	feeds.id, 
	feeds.created_at, 
	feeds.updated_at, 
	feeds.name, 
	feeds.url, 
	users.name AS user_name
FROM feeds
	INNER JOIN users on feeds.user_id = users.id;
--

-- name: GetFeedByURL :one
SELECT id, created_at, updated_at, name, url
FROM feeds
WHERE url = $1;
--

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = NOW() , updated_at = NOW()
WHERE id = $1;
--

-- name: GetNextFeedToFetch :one
SELECT id, name, url
FROM feeds
ORDER BY last_fetched_at DESC NULLS FIRST
LIMIT 1;
--
