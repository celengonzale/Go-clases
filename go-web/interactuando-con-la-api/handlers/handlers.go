package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-web/interactuando-con-la-api/Model"
	"io"
	"net/http"
	"os"
	"strconv"
)

func Ping(context *gin.Context) {
	context.String(200, "pong")
}

var products = DecodeFile()

func OpenJsonFile() (file *os.File, err error) {
	file, err = os.Open("products.json")

	if err != nil {
		panic(err)
	}
	return file, nil

}
func DecodeFile() (products Model.Products) {
	file, err := OpenJsonFile()
	if err != nil {
		panic(err)
	}
	byteValue, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(byteValue, &products.Products)
	if err != nil {
		panic(err)
	}
	return products

}
func GetAllProducts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Model.Response{
		Message: "Success: All Products",
		Data:    products,
	})
}
func GetProductById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Model.Response{
			Message: "Fail to parse id",
		})
		return
	}
	var item Model.Product
	for _, currentProduct := range products.Products {
		if currentProduct.Id == id {
			item = currentProduct
			break
		}
	}

	if item.Id == 0 {
		ctx.JSON(http.StatusNotFound, Model.Response{
			Message: "Product not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, Model.Response{Message: "Success Product", Data: item})

}
func GetProductsByPrice(ctx *gin.Context) {
	price, err := strconv.ParseFloat(ctx.Query("priceGt"), 64)
	if err != nil {
		ctx.JSON(http.StatusNotFound, Model.Response{
			Message: "Fail to parse PriceGt",
		})
		return
	}
	if price <= 0 {
		ctx.JSON(http.StatusNotFound, Model.Response{
			Message: "The price is not valid",
		})
		return
	}
	var prod Model.Products
	for _, product := range products.Products {
		if product.Price > price {
			prod.Products = append(prod.Products, product)
		}
	}
	ctx.JSON(http.StatusOK, Model.Response{
		Message: "Success Price",
		Data:    prod,
	})
}
