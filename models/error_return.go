package models

import "errors"

var (
	ErrorUrlNaoInformada    = errors.New("URL não foi informada")
	ErrorAliasExiste        = errors.New("Alias existente")
	ErrorInterno            = errors.New("Ocorreu um error interno na API")
	ErrorAliasNaoEncontrado = errors.New("Não encontrou o Alias informado")
	ErrorUrlInvalida        = errors.New("A URL informada está inválida")
	ErrorRedirectURL        = errors.New("Ocorreu um erro ao tentar redirecionar para url")
)

type ErrorReturn struct {
	URL         string `json:"url,omitempty"`
	Alias       string `json:"alias,omitempty"`
	CodigoError string `json:"codigo_error,omitempty"`
	Descricao   string `json:"descricao,omitempty"`
}

func NewErrorReturn(codigo string, descricao string) *ErrorReturn {
	return &ErrorReturn{"", "", codigo, descricao}
}
func NewErrorUrlNaoInformada() *ErrorReturn {
	return &ErrorReturn{"", "", "003", ErrorUrlNaoInformada.Error()}
}
func NewErrorAliasExiste() *ErrorReturn {
	return &ErrorReturn{"", "", "001", ErrorAliasExiste.Error()}
}
func NewErrorInterno() *ErrorReturn {
	return &ErrorReturn{"", "", "000", ErrorInterno.Error()}
}
func NewErrorAliasNaoEncontrado() *ErrorReturn {
	return &ErrorReturn{"", "", "002", ErrorAliasNaoEncontrado.Error()}
}
func NewErrorUrlInvalida(url string) *ErrorReturn {
	return &ErrorReturn{url, "", "004", ErrorUrlInvalida.Error()}
}
func NewErrorToRedirect(url string) *ErrorReturn {
	return &ErrorReturn{url, "", "005", ErrorRedirectURL.Error()}
}
