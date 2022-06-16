package postgresDB

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"local/wwc_cocksize_bot/pkg/models"
	"time"
)

type RefreshTokenRepository struct {
	db   *sqlx.DB
	psql sq.StatementBuilderType
}

func NewRefreshTokenRepository(db *sqlx.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{db: db, psql: sq.StatementBuilder.PlaceholderFormat(sq.Dollar)}
}

func (r RefreshTokenRepository) Create(userId int64, refreshToken string, expiresIn time.Time) (models.RefreshToken, error) {
	var rt models.RefreshToken
	query := fmt.Sprint("INSERT INTO  oauth_refresh_tokens (user_id, refresh_token, expires_in) values ($1, $2, $3) RETURNING *")
	row := r.db.QueryRowx(query, userId, refreshToken, expiresIn)
	if err := row.StructScan(&rt); err != nil {
		return models.RefreshToken{}, err
	}

	return rt, nil
}

func (r RefreshTokenRepository) Find(refreshToken string) (models.RefreshToken, error) {

	var rt models.RefreshToken

	qry := r.psql.Select("*").From("oauth_refresh_tokens").
		Where(sq.Eq{"refresh_token": refreshToken}).
		Where(sq.GtOrEq{"expires_in": time.Now()})
	sql, params, err := qry.ToSql()
	if err != nil {
		return models.RefreshToken{}, nil
	}
	err = r.db.Get(&rt, sql, params...)
	if err != nil {
		return models.RefreshToken{}, nil
	}
	return rt, nil
}
