/*
功能：网络
说明:
*/
package USDT_NETWORK

type item struct {
	t       TYPE   //类型
	txt_eng string //类型的文本描述
	txt_chs string
}

var item_map map[TYPE]item

func init() {
	item_map = make(map[TYPE]item, 0)
	items := []item{
		{BTC, "NETWORK_BTC", "比特币"},
		{ETH, "NETWORK_ETH", "以太坊"},
		{BSC, "NETWORK_BSC", "币安智能链"},
		{TRX, "NETWORK_TRX", "波场"},
	}

	for _, v := range items {
		item_map[v.t] = v
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
