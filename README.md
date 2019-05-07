# Challenge #5

The system needs to do the following things:

- Accept an order with a users contact information
- Have the ability to retrieve the users order history
- Create a PDF receipt that gets stored in a directory and a base64 encoded string of that file
- Stretch goal: Create a little frontend for this

```golang
type Order struct {
	Items           []Item
	UserEmail       string
	UserName        string
	UserPhonenumber string
}

type Item struct {
	ItemUPC  string
	Quantity int
}
```

This may come in handy: github.com/jung-kurt/gofpdf