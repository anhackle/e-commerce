package product_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/anle/codebase/internal/database"
	"github.com/anle/codebase/internal/model"
	"github.com/anle/codebase/internal/service"
	"github.com/stretchr/testify/assert"
)

// 🎭 Fake Repo trả về kiểu database.GetProductsRow như repo thật
type fakeProductRepo struct {
	getProductsFunc func(ctx context.Context, input model.GetProductInput) ([]database.GetProductsRow, error)
}

func (f *fakeProductRepo) CreateProduct(ctx context.Context, input model.CreateProductInput) (sql.Result, error) {
	return nil, nil // không test CreateProduct
}

func (f *fakeProductRepo) GetProducts(ctx context.Context, input model.GetProductInput) ([]database.GetProductsRow, error) {
	return f.getProductsFunc(ctx, input)
}

func TestProductService_GetProducts(t *testing.T) {
	tests := []struct {
		name           string
		input          model.GetProductInput
		mockRepoOutput []database.GetProductsRow
		mockRepoError  error
		expectedOutput []model.GetProductOutput
		expectedCode   int
		expectError    bool
	}{
		{
			name:  "success - multiple products",
			input: model.GetProductInput{Page: 1, Limit: 10},
			mockRepoOutput: []database.GetProductsRow{
				{
					ID:          2,
					Name:        "laptop",
					Description: sql.NullString{String: "Lenovo Thinkpad P17", Valid: true},
					Price:       45000000,
					Quantity:    100,
					ImageUrl:    "https://download.lenovo.com/pccbbs/mobiles_pdf/tp_p15_p17_p1_gen3_ubuntu_20.04_lts_installation_v1.0.pdf",
				},
			},
			mockRepoError: nil,
			expectedOutput: []model.GetProductOutput{
				{
					ID:          2,
					Name:        "laptop",
					Description: "Lenovo Thinkpad P17",
					Price:       45000000,
					Quantity:    100,
					ImageURL:    "https://download.lenovo.com/pccbbs/mobiles_pdf/tp_p15_p17_p1_gen3_ubuntu_20.04_lts_installation_v1.0.pdf",
				},
			},
			expectedCode: 20000,
			expectError:  false,
		},
		{
			name:           "repo error",
			input:          model.GetProductInput{Page: 1, Limit: 100},
			mockRepoOutput: nil,
			mockRepoError:  errors.New("db error"),
			expectedOutput: nil,
			expectedCode:   50000,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &fakeProductRepo{
				getProductsFunc: func(ctx context.Context, input model.GetProductInput) ([]database.GetProductsRow, error) {
					return tt.mockRepoOutput, tt.mockRepoError
				},
			}

			productService := service.NewProductService(mockRepo)
			result, code, err := productService.GetProducts(context.Background(), tt.input)

			assert.Equal(t, tt.expectedCode, code)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedOutput, result)
			}
		})
	}
}
