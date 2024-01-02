package entry

type Robot struct {
	ID           string `json:"id" gorm:"column:'id';primaryKey"`
	SN           string `json:"sn" gorm:"column:'sn';type:varchar(128);unique"`
	Name         string `json:"name" gorm:"column:'name';type:varchar(128);unique"`
	InternalIP   string `json:"internal_ip" gorm:"column:'internal_ip';type:varchar(128)"`
	ExternalIP   string `json:"external_ip" gorm:"column:'external_ip';type:varchar(128)"`
	BorrowedUser string `json:"borrowed_user" gorm:"column:'borrowed_user';type:varchar(128)"`
	BorrowedUsed string `json:"borrowed_used" gorm:"column:'borrowed_used';type:varchar(128)"`
	BorrowedDate string `json:"borrowed_date" gorm:"column:'borrowed_date';type:varchar(128)"`
}

type User struct {
	ID         string `json:"id" gorm:"column:'id';primaryKey"`
	Username   string `json:"username" gorm:"column:username;type:varchar(64);unique"`
	Password   string `json:"password" gorm:"column:password;type:varchar(64)"`
	RobotCount int    `json:"robot_count" gorm:"column:robot_count;type:int"`
	DeadCount  int    `json:"dead_count" gorm:"column:dead_count;type:int;comment:'过期的机器人数量'"`
}
