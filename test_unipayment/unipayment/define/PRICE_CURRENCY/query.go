/*
功能：价格货币
说明:
*/
package PRICE_CURRENCY

// 描述
func (t TYPE) String() string {
	item, ok := item_map[t]
	if ok {
		return item.txt_eng
	}
	return ""
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
