package market_service

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"testing"

	"app/api/handler"
	"app/api/models"

	"github.com/bxcodec/faker/v3"
	"github.com/spf13/cast"
	"github.com/test-go/testify/assert"
)

var s int64

func TestProduct(t *testing.T) {
	s = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 1000; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			id := createProduct(t)
			updateProduct(t, id)
			// deleteProduct(t, id)
		}()

		s++
	}

	wg.Wait()

	fmt.Println("s: ", s)
}

func createProduct(t *testing.T) string {
	response := &handler.Response{}

	request := &models.CreateProduct{
		Name:  faker.Name(),
		Price: float64(rand.Intn(1000000)),
	}

	resp, err := PerformRequest(http.MethodPost, "/product", request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	// fmt.Println(response.Data)

	return cast.ToString(cast.ToStringMap(response.Data)["id"])
}

func updateProduct(t *testing.T, id string) string {

	response := &handler.Response{}
	request := &models.UpdateProduct{Name: faker.Name(),
		Price: float64(rand.Intn(1000000)),
	}

	resp, err := PerformRequest(http.MethodPut, "/product/"+id, request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 200)
	}

	// fmt.Println(resp)

	return cast.ToString(cast.ToStringMap(response.Data)["id"])
}

func deleteProduct(t *testing.T, id string) string {

	resp, _ := PerformRequest(
		http.MethodDelete,
		fmt.Sprintf("/product/%s", id),
		nil,
		nil,
	)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return ""
}
