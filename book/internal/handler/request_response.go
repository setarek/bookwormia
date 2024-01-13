package handler

import (
	"bookwormia/book/internal/repository"
	"fmt"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type Respone struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type BookInfo struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	Link          string `json:"link"`
	BookmarkCount int64  `json:"bookmark_count"`
	Marked        bool   `json:"marked"`
}

func NewBookListResponse(baseUrl string, books []repository.BookInfo) []BookInfo {
	var response []BookInfo
	for _, book := range books {
		var bookInfo BookInfo
		bookInfo.Id = book.BookId
		bookInfo.Name = book.BookName
		if book.BookMarkedByUser > 0 {
			bookInfo.Marked = true
		}
		bookInfo.BookmarkCount = book.BookmarkCount
		bookInfo.Link = fmt.Sprintf("%s/api/v1/books/%d", baseUrl, book.BookId)
		response = append(response, bookInfo)
	}
	return response
}

type BookDetails struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	CommentCount int64   `json:"comment_count"`
	ScoreCount   int64   `json:"score_count"`
	AverageScore float64 `json:"average_score"`
}

type BookComment struct {
	UserId  int64  `json:"user_id"`
	Comment string `json:"comment"`
	Score   int64  `json:"score"`
}

type BookReviewRequest struct {
	BookId  int64  `json:"book_id"`
	Score   int64  `json:"score"`
	Comment string `json:"comment"`
}

type BookmarkRequest struct {
	BookId int64 `json:"book_id"`
	Maked  bool  `json:"marked"`
}

type BookReviewResponse struct {
	UserId  int64  `json:"user_id"`
	Score   int64  `json:"score"`
	Comment string `json:"comment"`
}

func NewReviewRespone(reviews []repository.BookReview) []BookReviewResponse {
	var response []BookReviewResponse
	for _, review := range reviews {
		var bookReview BookReviewResponse
		if review.UserId.Valid {
			bookReview.UserId = review.UserId.Int64
		}
		if review.Comment.Valid {
			bookReview.Comment = review.Comment.String
		}
		if review.Score.Valid {
			bookReview.Score = review.Score.Int64
		}
		response = append(response, bookReview)
	}
	return response

}
