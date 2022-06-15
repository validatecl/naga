package naga

import (
	"errors"
	"reflect"
)

//VariableType tipo de dato para tipo de variable
type VariableType int

const (
	//TypeInt Tipo de dato int
	TypeInt VariableType = iota
	//TypeBool Tipo de dato bool
	TypeBool
	//TypeString Tipo de dato string
	TypeString
	//TypeNone Tipo de dato no identificado
	TypeNone
	//ErrorNil error cuando la interface de entrada es nil
	ErrorNil = "El dato de entrada es nulo"
	//ErrorTypeNone error con el tipo de dato
	ErrorTypeNone = "Tipo de dato no soportado"
)

//VariableTypeResolver interface
type VariableTypeResolver interface {
	ResolveType(interface{}) (VariableType, error)
}

type variableTypeResolver struct {
}

//NewVariableTypeResolver constructor
func NewVariableTypeResolver() VariableTypeResolver {
	return &variableTypeResolver{}
}

//ResolveType retorna tipo de dato
func (v *variableTypeResolver) ResolveType(data interface{}) (VariableType, error) {

	var typeData VariableType //variable de retorno con el tipo de dato -> numero

	if data == nil {
		return TypeNone, errors.New(ErrorNil)
	}

	stringType := reflect.TypeOf(data).String()
	switch stringType {
	case "bool":
		typeData = TypeBool
	case "string":
		typeData = TypeString
	case "int":
		typeData = TypeInt
	default:
		return TypeNone, errors.New(ErrorTypeNone)
	}

	return typeData, nil

}
