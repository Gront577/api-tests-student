package favorite

import (
	"net/http"
	"testing"

	favorite "api-tests-template/internal/client/http/favorite"
	"api-tests-template/internal/managers/common"
	commonModels "api-tests-template/internal/managers/common/models"
	"api-tests-template/internal/managers/favorite/models"
)

// maxFavoritesPageSize максимальный размер страницы избранного из swagger
const maxFavoritesPageSize = 50

// GetFavorites получает избранное пользователя (максимальная страница), ожидает 200
func GetFavorites(t *testing.T, token string) models.FavoritesList {
	t.Helper()
	raw := favorite.HttpGetFavorites(t, token, maxFavoritesPageSize)

	favoritesList := models.FavoritesList{}
	common.DecodeBody(t, raw, http.StatusOK, &favoritesList)
	return favoritesList
}

// GetFavoritesWithError запрашивает избранное и возвращает тело ошибки, ожидает переданный статус-код
func GetFavoritesWithError(t *testing.T, token string, expectedStatusCode int) commonModels.ErrorResponse {
	t.Helper()
	raw := favorite.HttpGetFavorites(t, token, maxFavoritesPageSize)
	return common.DecodeError(t, raw, expectedStatusCode)
}

// AddFavorite добавляет объявление в избранное, ожидает 201
func AddFavorite(t *testing.T, token, advertisementId string) models.FavoriteItem {
	t.Helper()
	raw := favorite.HttpPostFavorite(t, token, models.AddFavoriteRequest{AdvertisementId: advertisementId})

	favorite := models.FavoriteItem{}
	common.DecodeBody(t, raw, http.StatusCreated, &favorite)
	return favorite
}

// AddFavoriteWithError пытается добавить объявление в избранное и возвращает тело ошибки, ожидает переданный статус-код
func AddFavoriteWithError(t *testing.T, token, advertisementId string, expectedStatusCode int) commonModels.ErrorResponse {
	t.Helper()
	raw := favorite.HttpPostFavorite(t, token, models.AddFavoriteRequest{AdvertisementId: advertisementId})
	return common.DecodeError(t, raw, expectedStatusCode)
}

// RemoveFavorite удаляет объявление из избранного, ожидает 204
func RemoveFavorite(t *testing.T, token, advertisementId string) {
	t.Helper()
	raw := favorite.HttpDeleteFavorite(t, token, advertisementId)
	common.ReadBodyWithStatus(t, raw, http.StatusNoContent)
}

// RemoveFavoriteWithError пытается удалить объявление из избранного и возвращает тело ошибки, ожидает переданный статус-код
func RemoveFavoriteWithError(t *testing.T, token, advertisementId string, expectedStatusCode int) commonModels.ErrorResponse {
	t.Helper()
	raw := favorite.HttpDeleteFavorite(t, token, advertisementId)
	return common.DecodeError(t, raw, expectedStatusCode)
}
