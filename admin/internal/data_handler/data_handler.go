package data_handler

type DataHandler interface {
	CreateOne(book Book) error
	UpdateOne(book Book) error
	DeleteOne(id int64) error

	CreateBulk(header []string, records [][]string) error
	UpdateBulk(header []string, records [][]string) error
	DeleteBulk(ids []int64) error
}
