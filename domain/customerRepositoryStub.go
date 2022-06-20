package domain


type CustomerRepositoryStub struct{
	customers []Customer
}

func (s CustomerRepositoryStub)FindAll()([]Customer,error){
	return s.customers,nil
}

func NewCustomerRepositoryStub()CustomerRepositoryStub{
	customers:=[]Customer{
		{
			Id: "1001",Name: "Karan",City: "New Delhi", ZipCode: "110005",DateofBirth: "1996-10-03",Status: "1",
		},
		{
			Id: "1002",Name: "Ashish",City: "New Delhi", ZipCode: "110005",DateofBirth: "1996-10-03",Status: "1",
		},
	}

	return CustomerRepositoryStub{customers: customers}
}