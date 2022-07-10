package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ArticlesModel = (*customArticlesModel)(nil)

type (
	// ArticlesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticlesModel.
	ArticlesModel interface {
		articlesModel
		FindAllByAuthorId(ctx context.Context, authorId int64, limit int, page int, queryOrder string) ([]Articles, error)
	}

	customArticlesModel struct {
		*defaultArticlesModel
	}
)

// NewArticlesModel returns a model for the database table.
func NewArticlesModel(conn sqlx.SqlConn) ArticlesModel {
	return &customArticlesModel{
		defaultArticlesModel: newArticlesModel(conn),
	}
}

func (m *defaultArticlesModel) FindAllByAuthorId(ctx context.Context, authorId int64, limit int, page int, queryOrder string) (articles []Articles, err error) {
	query := fmt.Sprintf("select %s from %s where author_id = $1  order by $2 limit $3, $4", articlesRows, m.table)
	if queryOrder == "" {
		queryOrder = "id DESC"
	}
	skip := 0
	if page > 0 {
		skip = (page - 1) * limit
	}
	articles = []Articles{}
	err = m.conn.QueryRowsCtx(ctx, &articles, query, authorId, queryOrder, limit, skip)
	switch err {
	case nil:
		return articles, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
