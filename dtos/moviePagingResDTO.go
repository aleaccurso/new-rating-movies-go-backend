package dtos

type MoviePagingResDTO struct {
	Page      int8
	Size      int8
	NbPages   int8
	NbResults int16
	Data      []MovieResDTO
}
