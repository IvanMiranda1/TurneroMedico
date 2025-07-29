package domain

import "errors"

var ErrNotFound = errors.New("recurso no encontrado")

var ErrAlreadyExists = errors.New("el recurso ya existe")

var ErrValidation = errors.New("error de validacion")

var ErrResponseNil = errors.New("la respuesta fue nil")

var ErrFailedToCreate = errors.New("fallo la creacion")

var ErrRegexMatch = errors.New("no cumple con el regex (expresion regular)")
