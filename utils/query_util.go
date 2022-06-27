package utils

const (
	INSERT_CUSTOMER = `insert into customer 
						(id,name,address,phone,email,saldo) 
						values (:id,:name,:address,:phone,:email,:saldo)`

	UPDATE_CUSTOMER = `update customer set name=:name, address=:address, phone=:phone, email=:email where id=:id`

	DELETE_CUSTOMER = `update customer set is_status=0 where id=$1`

	INSERT_CUSTOMER_PS = `insert into customer 
						(id,name,address,phone,email,saldo) 
						values ($1,$2,$3,$4,$5,$6)`

	UPDATE_CUSTOMER_PS = `update customer set name=$1, address=$2, phone=$3, email=$4 where id=$5`

	DELETE_CUSTOMER_PS_HD = `delete from customer where id=$1`
	DELETE_CUSTOMER_PS_SD = `update customer set is_status=0 where id=$1`

	SELECT_CUSTOMER_ALL = `SELECT id, name, address, phone, email, saldo
						from customer order by created_at limit $1 offset $2`
	SELECT_CUSTOMER_BY_ID = `SELECT id, name, address, phone, email, saldo
						from customer where id=$1`
	SELECT_CUSTOMER_BY_NAME = `SELECT id, name, address, phone, email, saldo
						from customer where name ilike $1`
	// SELECT_COUNT_CUSTOMER = `SELECT count(id) from customer where is_status=1`
	SELECT_SUM_SALDO_CUSTOMER = `SELECT sum(saldo) from customer where is_status=1`

	SELECT_COUNT_CUSTOMER = `SELECT is_status,count(id) as total from customer group by is_status`
	SELECT_COUNT_CUSTOMER_ADDRESS = `SELECT address,count(id) as total from customer group by address`

	INSERT_PRODUCT = `insert into product (id,name,price,description,stock,store_id) values ($1,$1,$1,$1,$1,$1)`
	UPDATE_PRODUCT = `update product set name=:name, price=:price, description=:description, stock=:stock where id=:id`
	DELETE_PRODUCT = `update product set is_status=0 where id=$1`

	INSERT_SHOP = `insert into shop (id,no_siup,name,address,phone) values (:id, :no_siup, :name, :address, :phone)`
	UPDATE_SHOP = `update shop set name=:name, no_siup=:no_siup, address=:address, phone=:phone where id=:id`
	DELETE_SHOP = `update shop set is_status=0 where id=$1`
	SELECT_SHOP_WITH_PRODUCT = `select s.id, s.no_siup, s.name, p.id as product_id, p.name as product_name, p.price, p.stock from shop s join product p on s.id = p.store_id order by no_siup asc limit $1 offset $2`
)
