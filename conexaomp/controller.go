package conexaomp

// NovaConexaoMP - Retorna uma nova conexão MP
func NovaConexaoMP(accessToken, publicKey string) (*ConexaoMP, error) {
	var conexaoMP ConexaoMP

	conexaoMP.AccessToken = accessToken
	conexaoMP.PublicKey = publicKey

	// Testa se o access token é válido
	// Faço um search em payments e verifico um retorno diferente de sucesso
	//_, err := conexaoMP.SearchPayments(nil)

	//if err != nil {
	//	mensagem := "conexaomp.NovaConexaoMP - Erro ao testar access token: " + err.Error()
	//	log.Println(mensagem)
	//	return nil, errors.New(mensagem)
	//}

	// Testa se a public key é válida
	// Como testar a public key?

	return &conexaoMP, nil
}
