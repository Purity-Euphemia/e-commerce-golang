package services

import (
	"ecommerce-go/models"
	"ecommerce-go/repositories"
	"errors"
)

func AddReview(productID, userID uint, rating int, title, comment string) (*models.Review, error) {
	if rating < 1 || rating > 5 {
		return nil, errors.New("rating must be between 1 and 5")
	}

	review := &models.Review{
		ProductID: productID,
		UserID:    userID,
		Rating:    rating,
		Title:     title,
		Comment:   comment,
	}

	err := repositories.CreateReview(review)
	return review, err
}

func GetProductReviews(productID uint, page, pageSize int) ([]models.Review, int64, error) {
	return repositories.GetProductReviews(productID, page, pageSize)
}

func DeleteReview(id uint, userID uint) error {
	return repositories.DeleteReview(id, userID)
}

func UpdateProductRating(productID uint) error {
	reviews, _, err := repositories.GetProductReviews(productID, 1, 1000)
	if err != nil {
		return err
	}

	if len(reviews) == 0 {
		return repositories.UpdateProduct(productID, &models.Product{Rating: 0})
	}

	var totalRating float64
	for _, review := range reviews {
		totalRating += float64(review.Rating)
	}

	avgRating := totalRating / float64(len(reviews))
	return repositories.UpdateProduct(productID, &models.Product{Rating: avgRating})
}
