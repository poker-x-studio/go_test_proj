/*
功能：价格货币
说明:
*/
package PRICE_CURRENCY

type item struct {
	t       TYPE   //类型
	txt_eng string //类型的文本描述
	txt_chs string
}

var item_map map[TYPE]item

func init() {
	item_map = make(map[TYPE]item, 0)
	items := []item{
		{CHY, "CNY", "人民币"},
		{USD, "USD", "美元"},
	}

	for _, v := range items {
		item_map[v.t] = v
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
