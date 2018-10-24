package io

type BasketLine struct {
	ProductId int32
	Quantity  int32
}

type BasketLineRepository interface {
	All() ([]BasketLine, error)
	Increment(id int32) error
}

type basicBasketLineRepository struct {
	data map[int32]int32
}

func (b basicBasketLineRepository) All() ([]BasketLine, error) {
	response := []BasketLine{}
	for id, quantity := range b.data {
		response = append(response, BasketLine{
			ProductId: id,
			Quantity:  quantity,
		})
	}
	return response, nil
}

func (b basicBasketLineRepository) Increment(id int32) error {
	quantity, ok := b.data[id]
	if !ok {
		quantity = 0
	}
	b.data[id] = quantity + 1
	return nil
}

func NewBasketLineRepository() BasketLineRepository {
	return basicBasketLineRepository{data: map[int32]int32{}}
}
