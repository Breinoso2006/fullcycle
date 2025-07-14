package product

type ProductUseCase struct {
	repository ProductRepositoryInterface
}

func NewProductUseCase(repository ProductRepositoryInterface) *ProductUseCase {
	return &ProductUseCase{repository}
}

// em tese, o retorno de um use case deve ser um DTO
// Mas como o exemplo é simples, vamos retornar a entidade diretamente
func (u *ProductUseCase) GetProduct(id int) (Product, error) {
	return u.repository.GetProduct(id)
}
