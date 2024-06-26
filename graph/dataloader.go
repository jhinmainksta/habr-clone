package graph

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/jhinmainksta/habr-clone/graph/my_model"
	"gorm.io/gorm"
)

const commentLoaderKey = "userloader"

func DataloaderMiddleware(db *gorm.DB, next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		commentLoader := CommentLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(ids []string) ([][]*my_model.Comment, []error) {
				rawComments := []*my_model.Comment{}
				limit, ok := strconv.Atoi(r.Header.Get("limit"))
				if ok != nil {
					limit = 5
				}
				err := db.Where(" post_id in ? AND parent_id = ?", ids, 0).Limit(limit * len(ids)).Find(&rawComments).Error

				if len(rawComments) == 0 || err != nil {
					return nil, []error{err}
				}

				commentsMap := make(map[string][]*my_model.Comment, len(ids))

				for _, comment := range rawComments {
					commentsMap[strconv.Itoa(comment.PostID)] = append(commentsMap[strconv.Itoa(comment.PostID)], comment)
				}

				sortedComments := make([][]*my_model.Comment, len(ids))

				for i, id := range ids {

					sortedComments[i] = commentsMap[id]
				}

				return sortedComments, []error{err}
			},
		}

		ctx := context.WithValue(r.Context(), commentLoaderKey, &commentLoader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getCommentLoader(ctx context.Context) *CommentLoader {
	return ctx.Value(commentLoaderKey).(*CommentLoader)
}
