package main

import (
	"fmt"
	"test/unipayment"
	"test/unipayment/define/PAY_CURRENCY"
	"test/unipayment/define/PRICE_CURRENCY"
	"test/unipayment/define/USDT_NETWORK"
)

const (
	REDIRECT_URL = "https://www.google.com"
)

func main() {
	err := unipayment.Request_create_invoice(PRICE_CURRENCY.CHY, 11.9, PAY_CURRENCY.USDT, USDT_NETWORK.TRX, REDIRECT_URL)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// 签名
func Sign(client_id string, client_secret string, uri string, body string) string {
	/*
				method := METHOD
				timestamp := time.Now().Unix() //时间戳

				fmt.Println("client_id:", client_id)
				fmt.Println("client_secret:", client_secret)
				fmt.Println("method:", method)
				fmt.Println("body:", body)

				//go的url.QueryEscape与js的encodeURIComponent是不一样的,主要体现在对空格的处理
				request_uri := url.QueryEscape(uri)
				fmt.Println("request_uri:", request_uri)

				//随机值
				r := rand.New(rand.NewSource(time.Now().UnixNano()))
				r.Seed(time.Now().UnixNano())
				nonce := r.Int()

				fmt.Println(body)
				body_base64 := Md5_base64(body)
				fmt.Println(body_base64)
		unipayment.
				signatureRawData := fmt.Sprintf("%s%s%s%d%d%s", client_id, method, request_uri, timestamp, nonce, body)
				fmt.Println(signatureRawData)

				//
				hmac := hmac.New(sha256.New, []byte(client_secret))
				hmac.Write([]byte(signatureRawData))
				sha25 := hex.EncodeToString(hmac.Sum(nil))
				//fmt.Println("Result: " + sha25)
				hmac_base64 := base64.StdEncoding.EncodeToString([]byte(sha25))
				//fmt.Println("Result: " + hmac_base64)

				return fmt.Sprintf("%s:%s:%d:%d", client_id, hmac_base64, nonce, timestamp)
	*/
	return "nil"
}
