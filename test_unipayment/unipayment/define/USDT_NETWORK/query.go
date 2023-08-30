/*
功能：网络
说明:
*/
package USDT_NETWORK

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
