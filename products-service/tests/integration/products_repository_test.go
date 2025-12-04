package tests_test

import (
	"context"
	"math"
	"testing"

	"github.com/at-kh/guru-apps-test-services/products-service/internal/api/domain/products"
	"github.com/at-kh/guru-apps-test-services/products-service/internal/api/repositories"
	repoproducts "github.com/at-kh/guru-apps-test-services/products-service/internal/api/repositories/products"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

const (
	testPostgresDriver = "pgx"
	testPostgresDSN    = "postgres://postgres:root@localhost:5432/products_service?sslmode=disable"

	testCount = 10_000
)

// prettyJSON helper function.
func prettyJSON(data any) string {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(b)
}

// newTxRepo helper function.
func newTxRepo(t *testing.T) (*sqlx.Tx, repositories.ProductsRepository) {
	db, err := sqlx.Open(testPostgresDriver, testPostgresDSN)
	require.NoError(t, err)

	tx, err := db.BeginTxx(context.Background(), nil)
	require.NoError(t, err)

	repo := repoproducts.NewRepository(tx, zap.NewExample())

	return tx, repo
}

// rollbackTx helper function.
func rollbackTx(t *testing.T, tx *sqlx.Tx) {
	require.NoError(t, tx.Rollback())
}

func TestRepository_Create(t *testing.T) {
	t.Run("create product", func(t *testing.T) {
		tx, repo := newTxRepo(t)
		defer rollbackTx(t, tx)

		tests := []struct {
			name    string
			product products.Product
			wantErr bool
		}{
			{
				name: "[SUCCESS] valid product",
				product: products.Product{
					Name:        "test-" + uuid.NewString(),
					Vendor:      "vendorA",
					Description: "some",
					Price:       decimal.NewFromFloat(1.25),
				},
				wantErr: false,
			},
			{
				name: "[SUCCESS] zero price",
				product: products.Product{
					Name:        "free-" + uuid.NewString(),
					Vendor:      "vendorB",
					Description: "free",
					Price:       decimal.Zero,
				},
				wantErr: false,
			},
			{
				name: "[ERROR] empty name",
				product: products.Product{
					Name:        "",
					Vendor:      "vendorC",
					Description: "some",
					Price:       decimal.NewFromFloat(1),
				},
				wantErr: true,
			},
			{
				name: "[ERROR] negative price",
				product: products.Product{
					Name:        "neg-" + uuid.NewString(),
					Vendor:      "vendorD",
					Description: "some",
					Price:       decimal.NewFromFloat(-1),
				},
				wantErr: true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				created, err := repo.Create(context.Background(), tt.product)
				if tt.wantErr {
					require.Error(t, err)
				} else {
					require.NoError(t, err)
					require.NotZero(t, created.ID)
					require.Equal(t, tt.product.Name, created.Name)
					require.Equal(t, tt.product.Vendor, created.Vendor)
					require.Equal(t, tt.product.Description, created.Description)
					require.Less(t, math.Abs(tt.product.Price.InexactFloat64()-created.Price.InexactFloat64()), 0.01)

					t.Log(prettyJSON(created))
				}
			})
		}

		t.Run("[ERROR] duplicate name + vendor", func(t *testing.T) {
			duplicateProduct := products.Product{
				Name:        "duplicate-product",
				Vendor:      "vendorX",
				Description: "some",
				Price:       decimal.NewFromFloat(10),
			}

			tx, repo = newTxRepo(t)
			defer rollbackTx(t, tx)

			// first insert
			_, err := repo.Create(context.Background(), duplicateProduct)
			require.NoError(t, err)

			// second insert
			_, err = repo.Create(context.Background(), duplicateProduct)
			require.Error(t, err) // should fail on unique constraint
		})
	})
}

func TestRepository_Delete(t *testing.T) {
	t.Run("delete product", func(t *testing.T) {
		tx, repo := newTxRepo(t)
		defer rollbackTx(t, tx)

		p := products.Product{
			Name:        "to_delete-" + uuid.NewString(),
			Vendor:      "vendorX",
			Description: "some",
			Price:       decimal.NewFromFloat(5),
		}
		created, err := repo.Create(context.Background(), p)
		require.NoError(t, err)

		// delete success
		err = repo.Delete(context.Background(), created.ID)
		require.NoError(t, err)

		// delete again -> should fail
		err = repo.Delete(context.Background(), created.ID)
		require.Error(t, err)

		// delete with nil -> should fail
		err = repo.Delete(context.Background(), uuid.Nil)
		require.Error(t, err)
	})
}

func TestRepository_GetAll(t *testing.T) {
	t.Run("pagination", func(t *testing.T) {
		tx, repo := newTxRepo(t)
		defer rollbackTx(t, tx)

		for i := 0; i < testCount; i++ {
			p := products.Product{
				Name:        "product-" + uuid.NewString(),
				Vendor:      "vendor-" + uuid.NewString(),
				Description: "some",
				Price:       decimal.NewFromFloat(float64(i + 1)),
			}
			_, err := repo.Create(context.Background(), p)
			require.NoError(t, err)
		}

		// limit=5, offset=0
		list, err := repo.GetAll(context.Background(), 5, 0)
		require.NoError(t, err)
		require.Len(t, list.Products, 5)
		require.Equal(t, uint64(testCount), list.Total)

		// limit=3, offset=10
		list, err = repo.GetAll(context.Background(), 3, 10)
		require.NoError(t, err)
		require.Len(t, list.Products, 3)
		require.Equal(t, uint64(testCount), list.Total)

		// limit=0, offset=0
		list, err = repo.GetAll(context.Background(), 0, 0)
		require.NoError(t, err)
		require.Len(t, list.Products, products.DefaultLimit)
		require.Equal(t, uint64(testCount), list.Total)

		// offset beyond total
		list, err = repo.GetAll(context.Background(), 5, testCount+1)
		require.NoError(t, err)
		require.Empty(t, list.Products)
		require.Equal(t, uint64(testCount), list.Total)
	})
}
