package util

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

// GenerateKeys product key
func GenerateKeys() (*rsa.PrivateKey, error) {
	// Generate a public/private key pair to use
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {

		return nil, err
	}
	//fmt.Println("Private Key: ", privateKey)
	//fmt.Println("Public Key: ", &privateKey.PublicKey)
	return privateKey, nil

}

// GenerateKeysAndExport to Export Keys in File
func GenerateKeysAndExport(filename string) error {
	privateKey, err := GenerateKeys()
	if err != nil {
		return err
	}

	if filename == "" {
		filename = "keys.pem"
	}
	pemfile, err := os.Create(filename)
	if err != nil {
		return err

	}
	// http://golang.org/pkg/encoding/pem/#Block
	varpemkey := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
	err = pem.Encode(pemfile, varpemkey)
	if err != nil {
		return err

	}
	defer pemfile.Close()
	return nil

}

// ImportKeyFile to import key from file
func ImportKeyFile(filename string) (*rsa.PrivateKey, error) {
	// import privateKey from file
	privateKeyFile, err := os.Open(filename)
	if err != nil {
		return nil, err

	}
	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)
	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)
	data, _ := pem.Decode([]byte(pembytes))
	privateKeyFile.Close()

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		return nil, err

	}
	return privateKeyImported, err

}

// SignatureToken trusted token
func SignatureToken(cookieToken []byte) (string, error) {
	privateKey, _ := GenerateKeys()
	// Instantiate a signer using RSASSA-PSS (SHA512) with the given private key.
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.PS512, Key: privateKey}, nil)
	if err != nil {
		return "", err
	}
	//fmt.Println(signer)
	// Sign a sample payload. Calling the signer returns a protected JWS object,
	// which can then be serialized for output afterwards. An error would
	// indicate a problem in an underlying cryptographic primitive.
	var payload = cookieToken
	object, err := signer.Sign(payload)
	if err != nil {
		return "", err
	}
	//fmt.Println(object)
	// Serialize the encrypted object using the full serialization format.
	// Alternatively you can also use the compact format here by calling
	// object.CompactSerialize() instead.
	serialized := object.FullSerialize()
	// Parse the serialized, protected JWS object. An error would indicate that
	// the given input did not represent a valid message
	//fmt.Println("serialized check : ", serialized)
	//fmt.Println("serialized format : ", reflect.TypeOf(serialized))
	return serialized, nil
}

// ParseSignature get token
func ParseSignature(Signature string) ([]byte, error) {
	object, err := jose.ParseSigned(Signature)
	if err != nil {
		return nil, err
	}
	privateKey, _ := ImportKeyFile("private.pem")

	// Now we can verify the signature on the payload. An error here would
	// indicate the the message failed to verify, e.g. because the signature was
	// broken or the message was tampered with.
	output, err := object.Verify(&privateKey.PublicKey)
	if err != nil {
		return nil, err
	}
	fmt.Printf(string(output))
	return output, err
}

// ConvertToJSON token to js
func ConvertToJSON(serialized []byte) (map[string]interface{}, error) {
	josemap := make(map[string]interface{})
	err := json.Unmarshal(serialized, &josemap)
	if err != nil {
		return nil, err
	}
	fmt.Println("payload value is ", josemap["payload"])
	//for key, value := range josemap {
	//  fmt.Println("index : ", key, " value : ", value)
	// }
	return josemap, nil
}

//SetToRedis add token to redis
func SetToRedis(username, token string) error {
	redis, err := RedisNewClient()
	if err != nil {
		return err
	}
	err = redis.Set(username, token, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetFromRedis retun key val
func GetFromRedis(username string) (string, error) {
	redis, _ := RedisNewClient()
	token, err := redis.Get(username).Result()
	if err != nil {
		return "", err
	}
	return token, nil
}

// GenerateJWTSigned test
func GenerateJWTSigned(privateClim interface{}) (string, error) {
	key := []byte("secret")
	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: key}, (&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		return "", err
	}

	cl := jwt.Claims{
		Subject:   "subject",
		Issuer:    "issuer",
		NotBefore: jwt.NewNumericDate(time.Now()),
		Expiry:    jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
	}

	// When setting private claims, make sure to add struct tags
	// to specify how to serialize the field. The naming behavior
	// should match the encoding/json package otherwise.
	privateCl := privateClim
	//privateCl := struct {
	//	Name string `json:"name"`
	//	Role string `json:"role"`
	//}{
	//	"majid zare",
	//	"admin",
	//}

	raw, err := jwt.Signed(sig).Claims(cl).Claims(privateCl).CompactSerialize()
	if err != nil {
		return "", err
	}

	//fmt.Println(raw)
	//fmt.Println(reflect.TypeOf(raw))
	// Output: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsibGVlbGEiLCJmcnkiXSwiY3VzdG9tIjoiY3VzdG9tIGNsYWltIHZhbHVlIiwiaXNzIjoiaXNzdWVyIiwibmJmIjoxNDUxNjA2NDAwLCJzdWIiOiJzdWJqZWN0In0.knXH3ReNJToS5XI7BMCkk80ugpCup3tOy53xq-ga47o
	return raw, nil
}

// ParseJSONWebTokenClaims test
func ParseJSONWebTokenClaims(tokP string, out2 interface{}) (jwt.Claims, error) {
	var sharedKey = []byte("secret")
	raw := tokP
	out1 := jwt.Claims{}
	tok, err := jwt.ParseSigned(raw)
	if err != nil {
		return jwt.Claims{}, err
	}

	if err := tok.Claims(sharedKey, &out1, &out2); err != nil {
		return jwt.Claims{}, err
	}
	//fmt.Printf("iss: %s, sub: %s, scopes: %s\n", out.Issuer, out.Subject, strings.Join(out2.Scopes, ","))
	// Output: iss: issuer, sub: subject, scopes: foo,bar
	return out1, nil
}

// GenerateJWTEncrypted Test
func GenerateJWTEncrypted(privateCli interface{}) (string, error) {
	sharedEncryptionKey := []byte("itsa16bytesecret")
	enc, err := jose.NewEncrypter(
		jose.A128GCM,
		jose.Recipient{Algorithm: jose.DIRECT, Key: sharedEncryptionKey},
		(&jose.EncrypterOptions{}).WithType("JWT"),
	)
	if err != nil {
		return "", err
	}

	cl := jwt.Claims{
		Subject:   "subject",
		Issuer:    "issuer",
		NotBefore: jwt.NewNumericDate(time.Now()),
		Expiry:    jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
	}

	privateCl := privateCli
	raw, err := jwt.Encrypted(enc).Claims(cl).Claims(privateCl).CompactSerialize()
	if err != nil {
		return "", err
	}

	//fmt.Println(raw)
	//fmt.Println(reflect.TypeOf(raw))
	return raw, nil
}

//ParseEncryptedToken test
func ParseEncryptedToken(tokenEnc string, castumOut interface{}) (jwt.Claims, error) {
	key := []byte("itsa16bytesecret")
	raw := tokenEnc
	defaultOut := jwt.Claims{}
	tok, err := jwt.ParseEncrypted(raw)
	if err != nil {
		return defaultOut, err
	}

	if err := tok.Claims(key, &defaultOut, &castumOut); err != nil {
		return defaultOut, err
	}
	//fmt.Printf("iss: %s, sub: %s\n", out.Issuer, out.Subject)
	//fmt.Printf("Name: %s, Role: %s\n", out2.Name, out2.Role)
	// Output: iss: issuer, sub: subject
	return defaultOut, nil
}
