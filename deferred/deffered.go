package deferred

type Buyer struct {
	ShopTransactionID string
	ShopOrderDate     string
	FullName          string
	FullNameKana      string
	ZipCode           string
	Address           string
	CompanyName       string
	DepartmentName    string
	Tel1              string
	Tel2              string
	Email             string
	Email2            string
	BilledAmount      string
	GMOExtend1        string
	PaymentType       string
	Sex               string
	BirthDay          string
	MemberRegistDate  string
	BuyCount          string
	BuyAmountTotal    string
	MemberID          string
}

func (o *Buyer) toParam() *buyer {
	p := &buyer{
		ShopTransactionID: o.ShopTransactionID,
		ShopOrderDate:     o.ShopOrderDate,
		FullName:          o.FullName,
		FullNameKana:      o.FullNameKana,
		ZipCode:           o.ZipCode,
		Address:           o.Address,
		CompanyName:       o.CompanyName,
		DepartmentName:    o.DepartmentName,
		Tel1:              o.Tel1,
		Tel2:              o.Tel2,
		Email:             o.Email,
		Email2:            o.Email2,
		BilledAmount:      o.BilledAmount,
		GMOExtend1:        o.GMOExtend1,
		PaymentType:       o.PaymentType,
		Sex:               o.Sex,
		BirthDay:          o.BirthDay,
		MemberRegistDate:  o.MemberRegistDate,
		BuyCount:          o.BuyCount,
		BuyAmountTotal:    o.BuyAmountTotal,
		MemberID:          o.MemberID,
	}
	return p
}

type Delivery struct {
	DeliveryCustomer *DeliveryCustomer
	Details          Details
}

func (o *Delivery) toParam() *delivery {
	p := &delivery{
		DeliveryCustomer: func() *deliveryCustomer {
			if o.DeliveryCustomer != nil {
				return o.DeliveryCustomer.toParam()
			}
			return nil
		}(),
		Details: o.Details,
	}
	return p
}

type Deliveries []*Delivery

type DeliveryCustomer struct {
	FullName       string
	FullNameKana   string
	ZipCode        string
	Address        string
	CompanyName    string
	DepartmentName string
	Tel            string
}

func (o *DeliveryCustomer) toParam() *deliveryCustomer {
	p := &deliveryCustomer{
		FullName:       o.FullName,
		FullNameKana:   o.FullNameKana,
		ZipCode:        o.ZipCode,
		Address:        o.Address,
		CompanyName:    o.CompanyName,
		DepartmentName: o.DepartmentName,
		Tel:            o.Tel,
	}
	return p
}

type Detail struct {
	DetailName     string
	DetailPrice    string
	DetailQuantity string
	GMOExtend2     string
	GMOExtend3     string
	GMOExtend4     string
	DetailBrand    string
	DetailCategory string
}

func (o *Detail) toParam() *detail {
	p := &detail{
		DetailName:     o.DetailName,
		DetailPrice:    o.DetailPrice,
		DetailQuantity: o.DetailQuantity,
		GMOExtend2:     o.GMOExtend2,
		GMOExtend3:     o.GMOExtend3,
		GMOExtend4:     o.GMOExtend4,
		DetailBrand:    o.DetailBrand,
		DetailCategory: o.DetailCategory,
	}
	return p
}

type Details []*Detail

type RegisterResponseParam struct {
	Result            string
	Errors            Errors
	TransactionResult *TransactionResult
}

func NewRegisterResponseParam(o *registerResponseParam) *RegisterResponseParam {
	p := &RegisterResponseParam{
		Result:            o.Result,
		Errors:            o.Errors,
		TransactionResult: o.TransactionResult,
	}
	return p
}

type Error struct {
	ErrorCode    string
	ErrorMessage string
}

func NewError(o *gmoError) *Error {
	p := &Error{
		ErrorCode:    o.ErrorCode,
		ErrorMessage: o.ErrorMessage,
	}
	return p
}

type Errors []*Error

type TransactionResult struct {
	ShopTransactionID string
	GMOTransactionID  string
	AuthorResult      string
}

func NewTransactionResult(o *transactionResult) *TransactionResult {
	p := &TransactionResult{
		ShopTransactionID: o.ShopTransactionID,
		GMOTransactionID:  o.GMOTransactionID,
		AuthorResult:      o.AuthorResult,
	}
	return p
}
