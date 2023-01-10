package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-web/interactuando-con-la-api/Model"
	"io"
	"net/http"
	"os"
	"regexp"
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
func AddProduct(ctx *gin.Context) {
	var product Model.Product

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, Model.Response{
			Message: "Error Should Bind",
			Data:    nil,
		})
		return
	}
	verification := verificationAllFields(product)
	if verification.Data != true {
		ctx.JSON(http.StatusBadRequest, verification)
		return
	}
	product.Id = products.Products[len(products.Products)-1].Id + 1
	products.Products = append(products.Products, product)
	ctx.JSON(http.StatusOK, Model.Response{
		Message: fmt.Sprintf("¡Product successfully added! Your ID %v is", product.Id),
		Data:    products,
	})
}
func verificationAllFields(product Model.Product) Model.Response {
	var setResponse Model.Response

	if err := validator.New().Struct(product); err != nil {
		setResponse = Model.Response{
			Message: "Invalid fields",
			Data:    err.Error(),
		}
	}
	if verification := codeValueVerification(product.CodeValue); verification.Data != true {
		setResponse = verification
	} else if verification := dateVerification(product.Expiration); verification.Data != true {
		setResponse = verification
	}

	return setResponse

}

func codeValueVerification(item string) Model.Response {
	for _, currentProduct := range products.Products {
		if currentProduct.CodeValue == item {
			fmt.Println(currentProduct.CodeValue)
			return Model.Response{
				Message: "The value of CodeValue already exists",
				Data:    false,
			}
		}
	}
	return Model.Response{
		Message: "Success: code value verification",
		Data:    true,
	}
}

func dateVerification(item string) Model.Response {
	dateRe := "^([0-9]{2})+([/])([0-9]{2})+([/])([0-9]{4})$"
	reg, _ := regexp.Compile(dateRe)
	if reg.MatchString(item) {
		return Model.Response{
			Message: "Success: date verification",
			Data:    true,
		}
	} else {
		return Model.Response{
			Message: "The date format is not valid",
			Data:    false,
		}
	}

}
