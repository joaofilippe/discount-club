package queries

const InsertDiscount = `
	INSERT INTO discounts (
		id,
		restaurant_id,
		user_id,
		code,
		percentage,
		start_date,
		end_date,
		times,
		description,
		created_at,
		updated_at,
		deleted_at,
		active
	) VALUES (
	 	:id,
		:restaurant_id,
		:user_id,
		:code,
		:percentage,
		:start_date,
		:end_date,
		:times,
		:description,
		:created_at,
		:updated_at,
		:deleted_at,
		:active
	);
`
const SelectDiscountByID = `
	SELECT * FROM discounts WHERE id = $1;
`

const SelectDiscountByCode = `
	SELECT * FROM discounts WHERE code = $1;
`
