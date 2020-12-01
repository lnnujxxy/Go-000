package service

import "gogeekbang/internal/dao"

// errors 直接往上层抛
func GetName(id int) (name string, err error) {
	return dao.GetName(id)
}
