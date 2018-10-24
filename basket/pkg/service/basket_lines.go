package service

import (
	"github.com/jmvargas/go-kit-example/basket/pkg/io"
)

type BasketLineWithProductData struct {
	io.BasketLine
	Title string
	Price float32
}
