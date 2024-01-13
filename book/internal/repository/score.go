package repository

import (
	"context"
	"database/sql"
)

func (b *BookRepository) UpsertScore(ctx context.Context, userId int64, bookId int64, score int64) error {
	query := `
		insert into book_scores (user_id, book_id, score) values ($1, $2, $3) on conflict (user_id, book_id) 
		do update set score = excluded.score;
	`

	if _, err := b.DB.ExecContext(ctx, query, userId, bookId, score); err != nil {
		return err
	}
	return nil
}

func (b *BookRepository) GetBookAvgScore(ctx context.Context, bookId int64) (float64, error) {

	query := `
		select avg(bs.score) from books b join book_scores bs on b.id = bs.book_id 
		where b.id = $1 group by b.id;
	`

	var avgScore float64
	if err := b.DB.QueryRow(query, bookId).Scan(&avgScore); err == sql.ErrNoRows {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return avgScore, nil
}

type DistictScore struct {
	ScoreOne   int64 `json:"score_one"`
	ScoreTwo   int64 `json:"score_two"`
	ScoreThree int64 `json:"score_three"`
	ScoreFour  int64 `json:"score_four"`
	ScoreFive  int64 `json:"score_five"`
}

func (b *BookRepository) GetDisticScore(ctx context.Context, bookId int64) (DistictScore, error) {

	query := `
		select count(case when bs.score = 1 then 1 end) as score_one,
		count(case when bs.score = 2 then 1 end) as score_two,
		count(case when bs.score = 3 then 1 end) as score_three,
		count(case when bs.score = 4 then 1 end) as score_four,
		count(case when bs.score = 5 then 1 end) as score_five 
		from books b join book_scores bs on b.id = bs.book_id where b.id = $1
		group by b.id, b.name;
	`

	var score DistictScore
	if err := b.DB.QueryRowContext(ctx, query, bookId).Scan(&score.ScoreOne, &score.ScoreTwo, &score.ScoreThree, &score.ScoreFour, &score.ScoreFive); err == sql.ErrNoRows {
		return DistictScore{}, nil
	} else if err != nil {
		return DistictScore{}, err
	}
	return score, nil
}

func (b *BookRepository) GetBookScoreByUserId(ctx context.Context, userId int64, bookId int64) (int64, error) {

	query := `
		select score from book_scores where user_id = $1 and book_id = $2;
	`

	var score int64
	if err := b.DB.QueryRowContext(ctx, query, userId, bookId).Scan(&score); err == sql.ErrNoRows {
		return 0, nil
	} else if err != nil {
		return 0, err
	}
	return score, nil
}
