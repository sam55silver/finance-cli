package lib

type Transaction struct {
	CategoryID int
	Amount     float64
	Title      string
}

func GetTransactionTable() string {
	return `
	create table if not exists transactions (
		id INTEGER PRIMARY KEY,
		title text not null,
		amount decimal(10, 2) not null,
		date timestamp default current_timestamp,
		category_id int not null,
		foreign key (category_id) references category(id)
	);	
	`
}
