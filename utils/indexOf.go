package utils

import "new-rating-movies-go-backend/models"

func IndexOfInt32(slice []int32, number int32) int {

	for idx, element := range slice {
		if element == number {
			return idx
		}
	}
	return -1
}

func IndexOfRate(rates []models.UserRate, movieDbId int32) int {
	for idx, rate := range rates {
		if rate.MovieDbId == movieDbId {
			return idx
		}
	}
	return -1
}
