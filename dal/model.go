package dal

type Project struct {
	ID          int    `json:"id" gorm:"primaryKey; autoIncrement"`
	Name        string `json:"name"gorm:"name;type:varchar(500);not null"`
	Description string `json:"description"gorm:"type:varchar(1000);description"`
}
type Column struct {
	ID        int    `json:"id"gorm:"primaryKey; AUTO_INCREMENT"`
	Name      string `json:"name"gorm:"name;type:varchar(255);not null;unique"`
	ProjectID int    `json:"project_id"gorm:"project_id; not null"`
	OrderNum  int    `json:"order_number"gorm:"order_number"`
	Status    string `json:"status"gorm:"status"`
}
type Task struct {
	ID          int    `json:"id"gorm:"primaryKey; autoIncrement; not null"`
	Name        string `json:"name"gorm:"name;type:varchar(500); not null"`
	Status      bool   `json:"status"gorm:"status"`
	Description string `json:"description"gorm:"type:varchar(5000);description"`
	ColumnID    int    `json:"column_id"gorm:"column_id; not null"`
}
type Comment struct {
	Description string `json:"description"gorm:"description;type:varchar(5000)"`
	TaskID      int    `json:"task_id"gorm:"task_id; not null"`
	ID          int    `json:"id"gorm:"primaryKey; autoIncrement"`
}
