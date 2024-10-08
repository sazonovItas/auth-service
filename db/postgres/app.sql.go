// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: app.sql

package db

import (
	"context"
)

const createApp = `-- name: CreateApp :one
insert into apps (name, secret, scope) 
values ($1, $2, $3) returning id, name, secret, scope
`

type CreateAppParams struct {
	Name   string
	Secret string
	Scope  string
}

func (q *Queries) CreateApp(ctx context.Context, arg CreateAppParams) (App, error) {
	row := q.db.QueryRow(ctx, createApp, arg.Name, arg.Secret, arg.Scope)
	var i App
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Secret,
		&i.Scope,
	)
	return i, err
}

const deleteApp = `-- name: DeleteApp :exec
delete from apps where id = $1
`

func (q *Queries) DeleteApp(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteApp, id)
	return err
}

const getApp = `-- name: GetApp :one
select id, name, secret, scope from apps where id = $1 limit 1
`

func (q *Queries) GetApp(ctx context.Context, id int64) (App, error) {
	row := q.db.QueryRow(ctx, getApp, id)
	var i App
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Secret,
		&i.Scope,
	)
	return i, err
}

const updateAppName = `-- name: UpdateAppName :exec
update apps set name = $2 where id = $1
`

type UpdateAppNameParams struct {
	ID   int64
	Name string
}

func (q *Queries) UpdateAppName(ctx context.Context, arg UpdateAppNameParams) error {
	_, err := q.db.Exec(ctx, updateAppName, arg.ID, arg.Name)
	return err
}

const updateAppScope = `-- name: UpdateAppScope :exec
update apps set scope = $2 where id = $1
`

type UpdateAppScopeParams struct {
	ID    int64
	Scope string
}

func (q *Queries) UpdateAppScope(ctx context.Context, arg UpdateAppScopeParams) error {
	_, err := q.db.Exec(ctx, updateAppScope, arg.ID, arg.Scope)
	return err
}

const updateAppSecret = `-- name: UpdateAppSecret :exec
update apps set secret = $2 where id = $1
`

type UpdateAppSecretParams struct {
	ID     int64
	Secret string
}

func (q *Queries) UpdateAppSecret(ctx context.Context, arg UpdateAppSecretParams) error {
	_, err := q.db.Exec(ctx, updateAppSecret, arg.ID, arg.Secret)
	return err
}
