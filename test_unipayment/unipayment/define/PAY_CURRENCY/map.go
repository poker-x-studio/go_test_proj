/*
功能：支付货币
说明:
*/
package PAY_CURRENCY

type item struct {
	t       TYPE   //类型
	txt_eng string //类型的文本描述
	txt_chs string
}

var item_map map[TYPE]item

func init() {
	item_map = make(map[TYPE]item, 0)
	items := []item{
		{USDT, "USDT", "USDT"},
	}

	for _, v := range items {
		item_map[v.t] = v
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
