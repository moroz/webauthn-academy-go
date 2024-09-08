-- name: InsertUserToken :one
insert into user_tokens (user_id, token, context) values ($1, $2, $3) returning *;

-- name: GetUserByToken :one
select u.* from users u join user_tokens ut on u.id = ut.user_id where ut.token = $1 and ut.context = $2 and ut.inserted_at > (now() at time zone 'utc') - (sqlc.arg(validity_days)::int * interval '1 day');
