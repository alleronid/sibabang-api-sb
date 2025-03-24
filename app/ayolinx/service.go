package ayolinx

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"lanaya/api/config"
	"math/big"
	"net/http"
	"strings"
	"time"

	"gorm.io/gorm"
)

type AppSetting struct {
	Key   string
	Value string
}

type DBUtil struct {
	db *gorm.DB
}

func NewDBUtil(db *gorm.DB) *DBUtil {
	return &DBUtil{db: db}
}

func (db *DBUtil) GetAppSetting(key string) string {
	var setting AppSetting
	db.db.Where("key = ?", key).First(&setting)
	return setting.Value
}

type AyolinxService struct {
	timestamp  string
	secretApp  string
	keySB      string
	secretSB   string
	enums      *AyolinxEnums
	httpClient *http.Client
}

func NewAyolinxService() *AyolinxService {
	db := NewDBUtil(config.GetDB())
	enums := NewAyolinxEnums()

	return &AyolinxService{
		timestamp:  time.Now().Format(time.RFC3339),
		keySB:      db.GetAppSetting("ayolinx_key_sb"),
		secretSB:   db.GetAppSetting("ayolinx_secret_sb"),
		secretApp:  db.GetAppSetting("sibabang_secret"),
		enums:      enums,
		httpClient: &http.Client{},
	}
}

func (a *AyolinxService) Signature() (string, error) {
	clientKey := a.secretSB
	requestTimestamp := a.timestamp
	stringToSign := clientKey + "|" + requestTimestamp

	// Read private key
	privateKeyBytes, err := ioutil.ReadFile("/home/alleroni/keys/private_key.pem")
	if err != nil {
		return "", fmt.Errorf("failed to read private key: %v", err)
	}

	// Parse private key
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		return "", fmt.Errorf("failed to parse PEM block containing private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %v", err)
	}


	hash := sha256.Sum256([]byte(stringToSign))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return "", fmt.Errorf("failed to sign: %v", err)
	}

	base64Signature := base64.StdEncoding.EncodeToString(signature)
	return base64Signature, nil
}

func (a *AyolinxService) GetToken() (string, error) {
	clientKey := a.keySB
	signature, err := a.Signature()
	if err != nil {
		return "", err
	}

	headers := map[string]string{
		"X-CLIENT-KEY": clientKey,
		"X-SIGNATURE":  signature,
	}

	response, err := a.API("/v1.0/access-token/b2b", headers, nil)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	if err := json.Unmarshal([]byte(response), &result); err != nil {
		return "", err
	}

	if accessToken, ok := result["accessToken"].(string); ok {
		return accessToken, nil
	}

	return "", fmt.Errorf("access token not found in response")
}

func (a *AyolinxService) API(url string, headers map[string]string, post interface{}) (string, error) {
	timestamp := a.timestamp
	baseURL := a.enums.URL_DEV + url

	defaultHeaders := map[string]string{
		"Content-Type": "application/json",
		"X-TIMESTAMP":  timestamp,
	}

	for k, v := range headers {
		defaultHeaders[k] = v
	}

	var reqBody []byte
	var err error
	if post != nil {
		reqBody, err = json.Marshal(post)
		if err != nil {
			return "", err
		}
	}

	req, err := http.NewRequest("POST", baseURL, strings.NewReader(string(reqBody)))
	if err != nil {
		return "", err
	}

	for k, v := range defaultHeaders {
		req.Header.Set(k, v)
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}

func (a *AyolinxService) BaseInterface(signature, timestamp, token, url string, post interface{}) (string, error) {
	baseURL := a.enums.URL_DEV + url

	reqBody, err := json.Marshal(post)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", baseURL, strings.NewReader(string(reqBody)))
	if err != nil {
		return "", err
	}

	req.Header.Set("X-TIMESTAMP", timestamp)
	req.Header.Set("X-SIGNATURE", signature)
	req.Header.Set("X-PARTNER-ID", a.keySB)
	req.Header.Set("X-EXTERNAL-ID", a.RandomNumber())
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}

func (a *AyolinxService) GenerateQris(data map[string]interface{}) (string, error) {
	timestamp := a.timestamp
	method := "POST"
	urlSignature := "/v1.0/qr/qr-mpm-generate"
	token, err := a.GetToken()
	if err != nil {
		return "", err
	}
	clientSecret := a.secretSB

	bodyJSON, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(bodyJSON)
	hexEncodedHash := fmt.Sprintf("%x", hash)

	dataToSign := fmt.Sprintf("%s:%s:%s:%s:%s", method, urlSignature, token, hexEncodedHash, timestamp)

	h := hmac.New(sha512.New, []byte(clientSecret))
	h.Write([]byte(dataToSign))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	response, err := a.BaseInterface(signature, timestamp, token, urlSignature, data)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (a *AyolinxService) WalletDana(data map[string]interface{}) (string, error) {
	timestamp := time.Now().Format(time.RFC3339)
	method := "POST"
	urlSignature := "/direct-debit/core/v1/debit/payment-host-to-host"
	clientSecret := a.secretSB
	token, err := a.GetToken()
	if err != nil {
		return "", err
	}

	// Marshal body to JSON
	bodyJSON, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// Calculate hash
	hash := sha256.Sum256(bodyJSON)
	hexEncodedHash := fmt.Sprintf("%x", hash)

	// Prepare data for HMAC
	dataToSign := fmt.Sprintf("%s:%s:%s:%s:%s", method, urlSignature, token, hexEncodedHash, timestamp)

	// Create HMAC
	h := hmac.New(sha512.New, []byte(clientSecret))
	h.Write([]byte(dataToSign))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	// Make request
	response, err := a.BaseInterface(signature, timestamp, token, urlSignature, data)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (a *AyolinxService) RandomNumber() string {
	// Generate a random integer between 11111111111 and 99999999999
	max := new(big.Int).SetInt64(99999999999)
	min := new(big.Int).SetInt64(11111111111)

	// Calculate the range
	diff := new(big.Int).Sub(max, min)
	diff = diff.Add(diff, big.NewInt(1))

	// Generate random number in range
	n, err := rand.Int(rand.Reader, diff)
	if err != nil {
		// Fallback to simpler random method if crypto/rand fails
		return fmt.Sprintf("%d", 11111111111+time.Now().UnixNano()%88888888889)
	}

	// Add min to get within range
	n = n.Add(n, min)

	return n.String()
}

func generateUniqueID() string {
	now := time.Now()
	sec := now.Unix()
	usec := now.Nanosecond() / 1000
	return fmt.Sprintf("%08x%05x", sec, usec)
}
