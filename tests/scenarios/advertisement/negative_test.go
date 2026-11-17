package advertisement

import (
	"fmt"
	"net/http"
	"strings"

	"api-tests-template/internal/assertions"
	apiErrors "api-tests-template/internal/constants/apiError"
	"api-tests-template/internal/constants/testData"
	fixtures "api-tests-template/internal/helpers/fix"
	"api-tests-template/internal/managers/advertisements"

	base "api-tests-template/tests"
)

// TestCreateAdvertisementWithoutRequiredFields — негативный сценарий:
// создание объявления без обязательных полей возвращает 400.
func (s *TestSuite) TestCreateAdvertisementWithoutRequiredFields() {
	request := fixtures.NewFullAdvertisementRequest()
	request.Title = ""

	base.Step("Пытаемся создать объявление с пустым title")
	errorResponse := advertisements.CreateAdvertisementWithError(
		s.T(), s.loginData.Token, request, nil, http.StatusBadRequest,
	)

	base.Check("Проверяем ошибку 400")
	assertions.ErrorResponseEquals(
		s.T(), apiErrors.CodeInvalidRequest, apiErrors.MessageTitleRequired, errorResponse,
	)
}

// TestCreateAdvertisementWithNegativePrice — негативный сценарий:
// отрицательная цена возвращает 400.
func (s *TestSuite) TestCreateAdvertisementWithNegativePrice() {
	request := fixtures.NewFullAdvertisementRequest()
	request.Price = -1

	base.Step("Пытаемся создать объявление с отрицательной ценой")
	errorResponse := advertisements.CreateAdvertisementWithError(
		s.T(), s.loginData.Token, request, nil, http.StatusBadRequest,
	)

	base.Check("Проверяем ошибку 400")
	assertions.ErrorResponseEquals(
		s.T(), apiErrors.CodeInvalidRequest, apiErrors.MessagePriceOutOfRange, errorResponse,
	)
}

// TestCreateAdvertisementWithTooLongTitle — негативный сценарий:
// слишком длинный title возвращает 400.
func (s *TestSuite) TestCreateAdvertisementWithTooLongTitle() {
	request := fixtures.NewFullAdvertisementRequest()
	request.Title = strings.Repeat("a", testData.TitleMaxLength+1)

	base.Step("Пытаемся создать объявление со слишком длинным title")
	errorResponse := advertisements.CreateAdvertisementWithError(
		s.T(), s.loginData.Token, request, nil, http.StatusBadRequest,
	)

	base.Check("Проверяем ошибку 400")
	assertions.ErrorResponseEquals(
		s.T(), apiErrors.CodeInvalidRequest,
		fmt.Sprintf(apiErrors.MessageTitleTooLongFormat, testData.TitleMaxLength, len(request.Title)),
		errorResponse,
	)
}

// TestCreateAdvertisementWithoutAuth — негативный сценарий:
// запрос без токена возвращает 401.
func (s *TestSuite) TestCreateAdvertisementWithoutAuth() {
	request := fixtures.NewFullAdvertisementRequest()

	base.Step("Пытаемся создать объявление без токена")
	errorResponse := advertisements.CreateAdvertisementWithError(
		s.T(), testData.EmptyToken, request, nil, http.StatusUnauthorized,
	)

	base.Check("Проверяем ошибку 401")
	assertions.ErrorResponseEquals(
		s.T(), apiErrors.CodeUnauthorized, apiErrors.MessageAuthHeaderRequired, errorResponse,
	)
}

// TestCreateAdvertisementWithInvalidToken — негативный сценарий:
// запрос с неверным токеном возвращает 401.
func (s *TestSuite) TestCreateAdvertisementWithInvalidToken() {
	request := fixtures.NewFullAdvertisementRequest()

	base.Step("Пытаемся создать объявление с неверным токеном")
	errorResponse := advertisements.CreateAdvertisementWithError(
		s.T(), testData.InvalidToken, request, nil, http.StatusUnauthorized,
	)

	base.Check("Проверяем ошибку 401")
	assertions.ErrorResponseEquals(
		s.T(), apiErrors.CodeUnauthorized, apiErrors.MessageInvalidOrExpiredAuth, errorResponse,
	)
}
