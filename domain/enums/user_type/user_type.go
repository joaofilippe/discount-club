package usertype

type UserType int

const (
	Admin UserType = iota
	RestaurantAdmin
	RestaurantWorker
	Client
)

func (ut UserType) String() string {
	return [...]string{"Admin", "RestaurantAdmin", "RestaurantWorker", "Client"}[ut]
}

func (ut UserType) Value() int {
	return int(ut)
}

func TryParse(value string) (UserType, bool) {
	switch value {
	case "Admin", "admin":
		return Admin, true
	case "RestaurantAdmin", "restaurantadmin", "restaurant-admin", "restaurant_admin":
		return RestaurantAdmin, true
	case "RestaurantWorker", "restaurantworker", "restaurant-worker", "restaurant_worker":
		return RestaurantWorker, true
	case "Client", "client":
		return Client, true
	default:
		return -1, false
	}
}

func FromInt(value int) (UserType, bool) {
	switch value {
	case 0:
		return Admin, true
	case 1:
		return RestaurantAdmin, true
	case 2:
		return RestaurantWorker, true
	case 3:
		return Client, true
	default:
		return -1, false
	}
}