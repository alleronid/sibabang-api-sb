package merchant

type MerchantService interface {
	GetMerchant(input HeaderInput) (Merchant, error)
	GetMerchantByID(merchantID int) (Merchant, error)
	GetMerchantByClientKey(clientKey string) (Merchant, error)
}

type serviceMerchant struct {
	repository Repository
}

func NewService(repository Repository) MerchantService {
	return &serviceMerchant{repository}
}

func (s *serviceMerchant) GetMerchant(input HeaderInput) (Merchant, error) {
	clientKey := input.XApiKey
	secretToken := input.SecretToken

	merchant, err := s.repository.GetMerchant(clientKey, secretToken)

	if err != nil {
		return merchant, err
	}

	return merchant, nil
}

func (s *serviceMerchant) GetMerchantByID(merchantID int) (Merchant, error) {
	merchant, err := s.repository.FindByID(merchantID)

	if err != nil {
		return merchant, err
	}

	return merchant, nil
}

func (s *serviceMerchant) GetMerchantByClientKey(clientKey string) (Merchant, error) {
	merchant, err := s.repository.FindByClientKey(clientKey)

	if err != nil {
		return merchant, err
	}

	return merchant, nil
}
