package nsdg1000t

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"net/http"
	"strings"
)

type Client struct {
	host             string
	user             string
	password         string
	getLoginResponse getLoginResponse
	cookie           string
	csrfToken        string
}

func decodeBody(resp *http.Response, v interface{}) error {
	return json.NewDecoder(resp.Body).Decode(v)
}

func NewClient(host, user, password string) *Client {
	return &Client{
		host:     host,
		user:     user,
		password: password,
	}
}

func (c *Client) decodeAndSetPublicKey() (*rsa.PublicKey, error) {
	pkey := "-----BEGIN PUBLIC KEY-----\n" + c.getLoginResponse.EncPubKey + "\n-----END PUBLIC KEY-----\n"
	block, _ := pem.Decode([]byte(pkey))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("failed to parse RSA public key")
	}
	return rsaPub, nil
}

func (c *Client) encryptedBase64Creds(rsaPub *rsa.PublicKey) (string, string, error) {
	encUser, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPub, []byte(c.user))
	if err != nil {
		return "", "", err
	}
	encPassword, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPub, []byte(c.password))
	if err != nil {
		return "", "", err
	}
	return base64.StdEncoding.EncodeToString(encUser), base64.StdEncoding.EncodeToString(encPassword), nil
}

func (c *Client) performRequest(method, url string, body *bytes.Buffer) (*http.Response, error) {
	if body == nil {
		body = bytes.NewBuffer([]byte{})
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if c.cookie != "" {
		req.Header.Set("Cookie", c.cookie)
	}
	req.Header.Set("Content-Type", "application/json")
	h := &http.Client{}
	return h.Do(req)
}

func (c *Client) getEndpointData(url string, v interface{}) error {
	resp, err := c.performRequest("GET", url, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return decodeBody(resp, v)
}

type getLoginResponse struct {
	LangCode   string `json:"lang_code"`
	DeviceName string `json:"device_name"`
	DeviceType string `json:"device_type"`
	FwVersion  string `json:"fw_version"`
	EncPubKey  string `json:"enc_pub_key"`
}

type postLoginRequest struct {
	getLoginResponse
	User       string `json:"username"`
	Password   string `json:"password"`
	ErrorCount string `json:"error_count"`
	WaitTime   string `json:"wait_time"`
}

func (c *Client) Login() error {
	url := "http://" + c.host + "/api/login"
	resp, err := c.performRequest("GET", url, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := decodeBody(resp, &c.getLoginResponse); err != nil {
		return err
	}

	rsaPub, err := c.decodeAndSetPublicKey()
	if err != nil {
		return err
	}

	user, password, err := c.encryptedBase64Creds(rsaPub)
	if err != nil {
		return err
	}

	loginReq := postLoginRequest{
		getLoginResponse: c.getLoginResponse,
		User:             user,
		Password:         password,
	}

	jsonRequest, err := json.Marshal(loginReq)
	if err != nil {
		return err
	}

	resp, err = c.performRequest("POST", url, bytes.NewBuffer(jsonRequest))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	c.cookie = resp.Header.Get("Set-Cookie")
	return nil
}

func getCSRFTokenFromBody(resp *http.Response) (string, error) {
	scan := bufio.NewScanner(resp.Body)
	for scan.Scan() {
		line := scan.Text()
		if strings.Contains(line, "csrf_token") {
			s := strings.Split(line, "'")
			if len(s) > 1 {
				return s[1], nil
			}
		}
	}
	return "", errors.New("failed to get CSRF token")
}

func (c *Client) GetCSRFToken() error {
	url := "http://" + c.host + "/pages.html"
	resp, err := c.performRequest("GET", url, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	csrfToken, err := getCSRFTokenFromBody(resp)
	if err != nil {
		return err
	}
	c.csrfToken = csrfToken
	return nil
}

type postLogoutRequest struct {
	User string `json:"username"`
}

func (c *Client) Logout() error {
	requst := postLogoutRequest{
		User: c.user,
	}
	jsonRequest, err := json.Marshal(requst)
	if err != nil {
		return err
	}
	url := "http://" + c.host + "/api/logout"
	resp, err := c.performRequest("POST", url, bytes.NewBuffer(jsonRequest))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
