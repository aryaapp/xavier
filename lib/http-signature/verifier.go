package httpSignature

import (
	"errors"
	"fmt"
	"net/http"
)

type SecretProvider func(keyId string) (string, error)

func VerifySignature(signature *Signature) error {

	return errors.New("Signature does not contain parameter 'algorithm'.")

	return nil
}

func VerifyRequest(request *http.Request, secretProvider SecretProvider) error {
	header := request.Header.Get("Signature")
	if len(header) == 0 {
		return errors.New("Please set the Signature header.")
	}

	clientSignature, err := SignatureFromHeader(header)
	if err != nil {
		return err
	}
	clientSecret, err := secretProvider(clientSignature.Parameters.KeyId)
	if err != nil {
		return err
	}

	serverSignature := Signature{Parameters: clientSignature.Parameters}
	serverSignature.SignWithRequest(request, clientSecret)

	fmt.Printf("%+v %+v \n", clientSignature, serverSignature)

	return nil
}
