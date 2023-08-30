/*
3.225.171.164
52.53.189.99
18.196.99.16
34.253.187.192
18.133.82.227
35.180.51.163
54.252.224.209
52.15.93.92
34.220.77.106
15.207.144.3
13.124.62.58
15.222.19.181
18.209.42.127
3.218.137.187
34.237.210.82
13.228.119.63
18.139.193.235
18.141.79.38
18.139.248.26
*/
package test

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/JFJun/trx-sign-go/grpcs"
	"github.com/JFJun/trx-sign-go/sign"
	"github.com/btcsuite/btcutil/base58"
	"github.com/fatih/structs"
	addr "github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/golang/protobuf/ptypes"
)

var (
	ADDR_FROM        = "TRQqvJia53MCFQVhe5so8HsTVqktBXCsXN"
	PRIVATE_KEY_FROM = "5ce0058cfd42618802af1122ff65771de0aba26054dd344692987b827558ebf7"
	ADDR_TO          = "THSLQG79pDbWhoL7EDwSXH55a2VSXV6xWk"
)

const (
	AMOUNT = 1 //转账个数
)

func Test_TransferTrx(t *testing.T) {
	//GRPC port: 50051
	c, err := grpcs.NewClient("18.139.248.26:50051")
	if err != nil {
		t.Fatal(err)
	}

	tx, err := c.Transfer(ADDR_FROM, ADDR_TO, AMOUNT)
	if err != nil {
		fmt.Println(111)
		t.Fatal(err)
	}
	signTx, err := sign.SignTransaction(tx.Transaction, PRIVATE_KEY_FROM)
	if err != nil {
		fmt.Println(222)
		t.Fatal(err)
	}
	err = c.BroadcastTransaction(signTx)
	if err != nil {
		fmt.Println(333)
		t.Fatal(err)
	}
	fmt.Println(common.BytesToHexString(tx.GetTxid()))

}

func Test_GetBalance(t *testing.T) {
	c, err := grpcs.NewClient("18.139.248.26:50051")
	if err != nil {
		t.Fatal(err)
	}
	acc, err := c.GetTrxBalance(ADDR_FROM)
	if err != nil {
		t.Fatal(err)
	}
	d, _ := json.Marshal(acc)
	fmt.Println(string(d))
	fmt.Println(acc.GetBalance())

}

func Test_GetTrc20Balance(t *testing.T) {
	c, err := grpcs.NewClient("grpc.trongrid.io:50051")
	if err != nil {
		t.Fatal(err)
	}
	amount, err := c.GetTrc20Balance("TK1UXQBkvAwBypz1bTWcuLHFaB8JmTjoUw", "TLdfZSUTwAJXxbav6od8iYCBSaW3EveYxm")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(amount.String())

}

func Test_TransferTrc20(t *testing.T) {
	c, err := grpcs.NewClient("54.168.218.95:50051")
	if err != nil {
		t.Fatal(err)
	}
	amount := big.NewInt(20)
	amount = amount.Mul(amount, big.NewInt(1000000000000000000))
	tx, err := c.TransferTrc20("TFTGMfp7hvDtt4fj3vmWnbYsPSmw5EU8oX", "TVwt3HTg6PjP5bbb5x1GtSvTe1J5FYM2BT",
		"TJ93jQZibdB3sriHYb5nNwjgkPPAcFR7ty", amount, 100000000)
	signTx, err := sign.SignTransaction(tx.Transaction, "5c023564aa0c582e9a5d127133e9b45c5b9a7a409b22f7e8a5c19d4d3f424eea")
	if err != nil {
		t.Fatal(err)
	}
	err = c.BroadcastTransaction(signTx)
	if err != nil {
		t.Fatal(err)

	}
	fmt.Println(common.BytesToHexString(tx.GetTxid()))
}

func Test_TransferTrc10(t *testing.T) {
	c, err := grpcs.NewClient("47.252.19.181:50051")
	if err != nil {
		t.Fatal(err)
	}
	from, _ := addr.Base58ToAddress("TFXf56UG1bhWkZq7WQEf7XW5hZXku17E8M")
	to, _ := addr.Base58ToAddress("TL4ebGiBbBPjduKaNEoPytVyzEuPEsFYz9")
	tokenID := "1000016"
	tx, err := c.GRPC.TransferAsset(from.String(), to.String(), tokenID, int64(123456))
	signTx, err := sign.SignTransaction(tx.Transaction, "")
	if err != nil {
		t.Fatal(err)
	}
	err = c.BroadcastTransaction(signTx)
	if err != nil {
		t.Fatal(err)

	}
	fmt.Println(common.BytesToHexString(tx.GetTxid()))
}

func Test_GetTrc10Balance(t *testing.T) {
	c, err := grpcs.NewClient("grpc.trongrid.io:50051")
	if err != nil {
		t.Fatal(err)
	}
	amount, err := c.GetTrc10Balance("TK1UXQBkvAwBypz1bTWcuLHFaB8JmTjoUw", "1002000")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(amount)
	//_, err := DecodeCheck("TFXf56UG1bhWkZq7WQEf7XW5hZXku17E8M")
	//if err != nil {
	//	t.Fatal(err)
	//}

}

func DecodeCheck(input string) ([]byte, error) {
	decodeCheck := base58.Decode(input)
	if len(decodeCheck) == 0 {
		return nil, fmt.Errorf("b58 decode %s error", input)
	}

	if len(decodeCheck) < 4 {
		return nil, fmt.Errorf("b58 check error")
	}

	decodeData := decodeCheck[:len(decodeCheck)-4]

	h256h0 := sha256.New()
	h256h0.Write(decodeData)
	h0 := h256h0.Sum(nil)

	h256h1 := sha256.New()
	h256h1.Write(h0)
	h1 := h256h1.Sum(nil)

	if h1[0] == decodeCheck[len(decodeData)] &&
		h1[1] == decodeCheck[len(decodeData)+1] &&
		h1[2] == decodeCheck[len(decodeData)+2] &&
		h1[3] == decodeCheck[len(decodeData)+3] {
		return decodeData, nil
	}
	return nil, fmt.Errorf("b58 check error")
}

func Test_GetBlock(t *testing.T) {
	c, err := grpcs.NewClient("47.252.19.181:50051")
	if err != nil {
		t.Fatal(err)
	}
	lb, err := c.GRPC.GetNowBlock()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(lb.BlockHeader.RawData.Number)
	fmt.Println(hex.EncodeToString(lb.Blockid))
}

func Test_GetTxByTxid(t *testing.T) {
	c, err := grpcs.NewClient("grpc.trongrid.io:50051")
	if err != nil {
		t.Fatal(err)
	}
	ti, err := c.GRPC.GetTransactionInfoByID("efe29a10301aae1c666a6db457708f0a9c452edbae81316a35a21483ce4772ab")
	if err != nil {
		t.Fatal(err)
	}

	fee := ti.Receipt.GetEnergyFee() + ti.Receipt.GetNetFee()
	fmt.Println(fee)
}

func Test_GetTransaction(t *testing.T) {
	c, err := grpcs.NewClient("3.225.171.164:50051")
	if err != nil {
		t.Fatal(err)
	}
	txid := "4dd50423887a722e812c70de058bb76b28ee33a2785ebaa6470dc0e2db8eeb53"
	txInfo, err := c.GRPC.GetTransactionByID(txid)
	if err != nil {
		t.Fatal(err)
	}
	d, _ := json.Marshal(txInfo)
	fmt.Println(string(d))
	r, err := c.GRPC.GetTransactionInfoByID(txid)
	if err != nil {
		t.Fatal(err)
	}
	dd, _ := json.Marshal(r)
	fmt.Println(string(dd))
	var cc core.TriggerSmartContract
	if err = ptypes.UnmarshalAny(txInfo.GetRawData().GetContract()[0].GetParameter(), &cc); err != nil {
		t.Fatal(err)
	}
	tv := structs.Map(cc)
	i := tv["Data"]
	da := i.([]uint8)
	data := hex.EncodeToString(da)
	if !strings.HasPrefix(data, trc20TransferMethodSignature) {
		t.Fatal("111")
	}

}

func Test_DecodeMessage(t *testing.T) {
	data := "CMN5oAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABlSYXRlIHN0YWxlIG9yIG5vdCBhIHN5bnRoAAAAAAAAAA=="
	d, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(d))
	dd, _ := hex.DecodeString("1952617465207374616c65206f72206e6f7420612073796e746800000000000000")
	fmt.Println(string(dd))
}

const trc20TransferMethodSignature = "a9059cbb"
