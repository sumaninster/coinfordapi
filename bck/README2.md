https://golang.org/doc/faq#convert_slice_of_interface

	//params := []string{"abc123"}
	//req := requestJSON{jsonrpc: "2.0", method: "personal_newAccount", params: params, id: "74"}
	//values := map[string]string{"jsonrpc":"2.0","method":"personal_newAccount","params":["passowrd"],"id":74}
	//values := map[string]string{"jsonrpc":"2.0","method":"personal_newAccount","params":"","id":"1"}
	//values := map[string]string{"jsonrpc":"2.0","method":"eth_accounts","id":"1"}
	//jsonValue, _ := json.Marshal({"jsonrpc":"2.0","method":"personal_newAccount","params":["passowrd"],"id":74})
	//resp, _ := http.Post(authAuthenticatorUrl, "application/json", bytes.NewBuffer(jsonValue))
	//curl -X POST --data '{"jsonrpc":"2.0","method":"personal_newAccount","params":["passowrd"],"id":74}' http://localhost:8545

	"io"
	"log"
	"strings"
	
	//var rqd responseJSON
	//json.Unmarshal(resp.ToByte(), &rqd)

	//ethp.print_values(body)

	//fmt.Println(string(body))
	//var dat map[string]interface{}
	//json.Unmarshal(body, &dat)
	//fmt.Println(dat)
	//fmt.Printf("%T: %v", dat, dat)
	//var accounts []string


func (u *UserController) prepareLoginResponse(tokenString string, user *models.User) LoginResponse {
	return LoginResponse{Token: tokenString, Name: user.Name, ResponseCode: 200, ResponseDescription: "Login success"}
}

func (ethp *ETHProcess) print_values(jsonStream []byte) {
	dec := json.NewDecoder(strings.NewReader(string(jsonStream)))
	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%T: %v", t, t)
		if dec.More() {
			fmt.Printf(" (more)")
		}
		fmt.Printf("\n")
	}
}

type requestJSON struct {
	jsonrpc 	string
	method 		string
	params		[]string
	id 			string
}
/*
type responseJSON struct {
	Resp 		string
	Headers 	Header
	Address 	string

}*/


type Header struct {
	ContentType 	string
	Date			time.Time
	ContentLength 	string
}

func (ethp *ETHProcess) json() {
    url := "http://restapi3.apiary.io/notes"
    fmt.Println("URL:>", url)

    var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}


--------Order------
	//"sort"
	
/*
func Orders2(user *User) (int64, []Order, error) {
	var table Order
	var orders []Order
	num, err := orm.NewOrm().QueryTable(table).Filter("user_id", user.Id).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&orders)
	return num, orders, err
}*/



/*
type lessFunc func(p1, p2 *Order) bool

type orderSorter struct {
	orders 	[]Order
	less    []lessFunc
}

func (ms *orderSorter) Sort(orders []Order) {
	ms.orders = orders
	sort.Sort(ms)
}

func OrderedBy(less ...lessFunc) *orderSorter {
	return &orderSorter{
		less: less,
	}
}

func (ms *orderSorter) Len() int {
	return len(ms.orders)
}

func (ms *orderSorter) Swap(i, j int) {
	ms.orders[i], ms.orders[j] = ms.orders[j], ms.orders[i]
}

func (ms *orderSorter) Less(i, j int) bool {
	p, q := &ms.orders[i], &ms.orders[j]
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}
	return ms.less[k](p, q)
}*/

	//amountSort := func(o1, o2 *models.Order) bool {
		//return o1.Amount < o2.Amount
	//}
	//models.OrderedBy(amountSort).Sort(*buyOrders)