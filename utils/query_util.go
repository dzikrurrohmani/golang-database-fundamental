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
)
