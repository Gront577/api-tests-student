package favorites

import (
	"fmt"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/favorite"

	base "api-tests-template/tests"
)

// TestFavoritesFullCycle — позитивный сценарий: полный цикл работы с избранным.
// Проверяет добавление, удаление из избранного и исчезновение объявления
// из избранного после его удаления владельцем.
func (s *TestSuite) TestFavoritesFullCycle() {
	advertisement := s.createOwnerAdvertisement()

	s.Run("Добавленное объявление появляется в избранном", func() {
		base.Step(fmt.Sprintf("Покупатель добавляет объявление %s в избранное", advertisement.Id))
		addedFavorite := favorite.AddFavorite(s.T(), s.buyer.Token, advertisement.Id)
		require.Equal(
			s.T(),
			advertisement.Id,
			addedFavorite.Advertisement.Id,
			"В ответе добавлено не то объявление",
		)

		base.Check("Проверяем, что объявление есть в GET /favorites")
		assertions.FavoritesContainAdvertisement(
			s.T(),
			favorite.GetFavorites(s.T(), s.buyer.Token),
			advertisement.Id,
		)
	})

	s.Run("Удалённое из избранного объявление пропадает из списка", func() {
		base.Step("Покупатель удаляет объявление из избранного")
		favorite.RemoveFavorite(s.T(), s.buyer.Token, advertisement.Id)

		base.Check("Проверяем, что объявления нет в GET /favorites")
		assertions.FavoritesNotContainAdvertisement(
			s.T(),
			favorite.GetFavorites(s.T(), s.buyer.Token),
			advertisement.Id,
		)
	})

	s.Run("Удалённое владельцем объявление пропадает из избранного", func() {
		base.Step("Покупатель снова добавляет объявление в избранное")
		favorite.AddFavorite(s.T(), s.buyer.Token, advertisement.Id)
		assertions.FavoritesContainAdvertisement(
			s.T(),
			favorite.GetFavorites(s.T(), s.buyer.Token),
			advertisement.Id,
		)

		base.Step("Владелец удаляет объявление через DELETE /advertisement")
		advertisements.DeleteAdvertisement(s.T(), s.owner.Token, advertisement.Id)

		base.Check("Проверяем, что объявления нет в избранном покупателя")
		assertions.FavoritesNotContainAdvertisement(
			s.T(),
			favorite.GetFavorites(s.T(), s.buyer.Token),
			advertisement.Id,
		)
	})
}
