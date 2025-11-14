package entity

import "github.com/google/uuid"

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

// Essa função converte a string para o tipo ID, retornando um erro caso a string não seja um UUID válido
func ParseID(s string) (ID, error) {
	id,err := uuid.Parse(s)
	return ID(id), err
}