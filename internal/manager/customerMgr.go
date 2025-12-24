package manager

import (
	"fmt"
	"hotel-mgmt-lld/internal/domain"
	"sync"
)

type CustomerMgr struct {
	CustomerList map[int]*domain.Customer
}

var (
	CustumerMgrInstance *CustomerMgr
	CustomerMgrOnce     sync.Once
)

func NewCustomerMgr() *CustomerMgr {
	CustomerMgrOnce.Do(func() {
		CustumerMgrInstance = &CustomerMgr{
			CustomerList: make(map[int]*domain.Customer),
		}
	})
	return CustumerMgrInstance
}

func (cm *CustomerMgr) AddCustomer(c *domain.Customer) {
	if c != nil {
		cm.CustomerList[c.UserId] = c
	}
}

func (cm *CustomerMgr) ReturnCustomerById(id int) (*domain.Customer, error) {
	for _, v := range cm.CustomerList {
		if v.UserId == id {
			return v, nil
		}
	}
	return nil, fmt.Errorf("Customer not found with id : %v\n", id)
}
