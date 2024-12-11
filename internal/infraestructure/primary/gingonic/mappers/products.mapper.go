package mappers

import (
	"velocity-technical-test/internal/domain/products/dtos"
	"velocity-technical-test/internal/infraestructure/primary/gingonic/response"
)

func ToProductResponse(dto dtos.ProductDTO) response.ProductResponse {
	return response.ProductResponse{
		ID:        dto.ID,
		Name:      dto.Name,
		Price:     dto.Price,
		Stock:     dto.Stock,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}

func ToProductDTO(resp response.ProductResponse) dtos.ProductDTO {
	return dtos.ProductDTO{
		ID:        resp.ID,
		Name:      resp.Name,
		Price:     resp.Price,
		Stock:     resp.Stock,
		CreatedAt: resp.CreatedAt,
		UpdatedAt: resp.UpdatedAt,
	}
}

func ToProductResponseList(dtos []dtos.ProductDTO) []response.ProductResponse {
	var responses []response.ProductResponse
	for _, dto := range dtos {
		responses = append(responses, ToProductResponse(dto))
	}
	return responses
}

func ToProductDTOList(responses []response.ProductResponse) []dtos.ProductDTO {
	var dtos []dtos.ProductDTO
	for _, resp := range responses {
		dtos = append(dtos, ToProductDTO(resp))
	}
	return dtos
}
