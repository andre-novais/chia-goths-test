// c.Post("/csrf-testing", func(w http.ResponseWriter, r *http.Request) {
// 	renderer.RenderHTML(r, w, "csrf-testing-post", r.PostForm)
// })

package handlers

import (
	"chia-goths/internal"
	"chia-goths/internal/application/service"
	"fmt"
	"strconv"

	"net/http"
)

type Controller struct {
	Renderer        internal.Renderer
	ProductsService service.ProdutsService
}

func (self *Controller) ListProducts(w http.ResponseWriter, r *http.Request) {

	products := self.ProductsService.List()

	self.Renderer.RenderHTML(r, w, "products", products)
}

func (self *Controller) CreateProduct(w http.ResponseWriter, r *http.Request) {

	//product := core.MakeProduct(30, "batata", "3241421")

	i, err := strconv.Atoi(r.Form.Get("price"))
	if err != nil {
		panic(err)
	}

	product := self.ProductsService.Create(i, r.Form.Get("name"), r.Form.Get("sku"))

	fmt.Print(r.Body, r.PostForm)

	self.Renderer.RenderHTML(r, w, "products-listing", product)
}
