package dtos

type UserPagingResDTO struct {
	Page        int8
	Size        int8
	RequestType string
	NbPages     int8
	NbResults   int16
	Data        []UserResDTO
}
