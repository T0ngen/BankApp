-- name: GetAccounts :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;


-- name: GetListAccounts :many
SELECT * FROM accounts
LIMIT $1;

-- name: CreateAccount :one
INSERT INTO accounts (
    owner, country_code,
    currency, balance,
    mail
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;


-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;


-- name: UpdateAccount :one
UPDATE accounts
SET owner = $2, country_code = $3, currency = $4, balance = $5, mail = $6
WHERE id = $1
RETURNING *;


-- name: GetEntries :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: GetListEntries :many
SELECT * FROM entries
LIMIT $1;

-- name: CreateEntrie :one
INSERT INTO entries (
    account_id, amount
) VALUES (
    $1, $2
)
RETURNING *;

-- name: DeleteEntrie :exec
DELETE FROM entries
WHERE id = $1;

-- name: UpdateEntrie :one
UPDATE entries
SET account_id = $2, amount = $3
WHERE id = $1
RETURNING *;








-- name: GetTransfers :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: GetListTransfers :many
SELECT * FROM transfers
LIMIT $1;

-- name: CreateTransfer :one
INSERT INTO transfers (
    from_account_id, to_account_id, amount
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1;

-- name: UpdateTransfer :one
UPDATE transfers
SET from_account_id = $2, to_account_id = $3, amount = $4
WHERE id = $1
RETURNING *;