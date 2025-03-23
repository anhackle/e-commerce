package product_test

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anle/codebase/internal/controller"
	"github.com/anle/codebase/internal/model"
	"github.com/anle/codebase/response"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type fakeProductService struct {
	mockProducts []model.GetProductOutput
	mockCode     int
	mockError    error
}

func (f *fakeProductService) GetProducts(ctx context.Context, input model.GetProductInput) ([]model.GetProductOutput, int, error) {
	return f.mockProducts, f.mockCode, f.mockError
}

func (f *fakeProductService) CreateProduct(ctx context.Context, input model.CreateProductInput) (int, error) {
	// Có thể mock thêm nếu test CreateProduct controller
	return response.ErrCodeSuccess, nil
}

func TestProductController_GetProducts(t *testing.T) {
	gin.SetMode(gin.TestMode)

	type mockService struct {
		mockProducts []model.GetProductOutput
		mockCode     int
		mockError    error
	}

	type args struct {
		requestBody string
	}

	tests := []struct {
		name           string
		mock           mockService
		args           args
		expectedStatus int
		expectContains string
	}{
		{
			name: "success - return product list",
			mock: mockService{
				mockProducts: []model.GetProductOutput{
					{ID: 1, Name: "Product A", Price: 100},
				},
				mockCode:  response.ErrCodeSuccess,
				mockError: nil,
			},
			args:           args{requestBody: `{"limit": 10, "page": 1}`},
			expectedStatus: http.StatusOK,
			expectContains: "Product A",
		},
		{
			name:           "invalid input - binding error",
			mock:           mockService{},
			args:           args{requestBody: `{}`},
			expectedStatus: http.StatusBadRequest,
			expectContains: "",
		},
		{
			name: "service error",
			mock: mockService{
				mockProducts: nil,
				mockCode:     response.ErrCodeInternal,
				mockError:    errors.New("internal error"),
			},
			args:           args{requestBody: `{"limit": 10, "page": 1}`},
			expectedStatus: http.StatusInternalServerError,
			expectContains: "",
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			mockService := &fakeProductService{
				mockProducts: tt.mock.mockProducts,
				mockCode:     tt.mock.mockCode,
				mockError:    tt.mock.mockError,
			}

			controller := controller.NewProductController(mockService)

			req, _ := http.NewRequest(http.MethodPost, "/products/search", bytes.NewBuffer([]byte(tt.args.requestBody)))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req

			controller.GetProducts(ctx)

			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.expectContains != "" {
				assert.Contains(t, w.Body.String(), tt.expectContains)
			}
		})
	}
}
