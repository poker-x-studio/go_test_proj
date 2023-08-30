package unipayment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"test/node"
	"test/unipayment/define/PAY_CURRENCY"
	"test/unipayment/define/PRICE_CURRENCY"
	"test/unipayment/define/USDT_NETWORK"
	"time"

	"github.com/google/uuid"
)

type CreateInvoice struct {
	App_id         string  `json:"app_id"`
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	Lang           string  `json:"lang"`
	Price_amount   float64 `json:"price_amount"`
	Price_currency string  `json:"price_currency"`
	Pay_currency   string  `json:"pay_currency"`
	Notify_url     string  `json:"notify_url"`
	Redirect_url   string  `json:"redirect_url"`
	Network        string  `json:"network"`
	Order_id       string  `json:"order_id"`
	Confirm_speed  string  `json:"confirm_speed"`
	Page_no        int     `json:"page_no"`
	Page_size      int     `json:"page_size"`
}

func (r *CreateInvoice) init(price_currency PRICE_CURRENCY.TYPE, price_amount float64, pay_currency PAY_CURRENCY.TYPE, network USDT_NETWORK.TYPE, redirect_url string) {
	r.App_id = APP_ID_BACCARAT
	r.Title = TITLE
	r.Description = DESCRIPTION
	r.Lang = LANG
	r.Price_amount = price_amount
	r.Price_currency = price_currency.String()
	r.Pay_currency = pay_currency.String()
	r.Network = network.String()
	r.Notify_url = WEBHOOK_URL
	r.Redirect_url = redirect_url
	r.Order_id = uuid.New().String()
	r.Confirm_speed = CONFIRM_SPEED
	r.Page_no = 1
	r.Page_size = 10
}

func (r *CreateInvoice) make_url() string {
	return API_HOST + "/v" + API_VERSION + "/invoices"
}

// 创建订单
// https://unipayment.readme.io/reference/create_invoice
func Request_create_invoice(price_currency PRICE_CURRENCY.TYPE, price_amount float64, pay_currency PAY_CURRENCY.TYPE, network USDT_NETWORK.TYPE, redirect_url string) error {
	var request CreateInvoice
	request.init(price_currency, price_amount, pay_currency, network, redirect_url)

	body, err := json.Marshal(request)
	if err != nil {
		return err
	}
	fmt.Println(string(body))

	url := request.make_url()
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Add("Accept", JSON)
	req.Header.Add("Content-Type", JSON)

	//签名
	sign, err := node.Call_javascript_sign(CLIENT_ID, CLIENT_SECRET, url, http.MethodPost, string(body))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", HMAC+sign)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Printf("resp:%+v", resp)

	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(resp_body))
	return nil
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
