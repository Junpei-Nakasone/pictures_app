package repository

type ServiceRepository interface {
	AddImage(userID int, fileURL string) error
}
