package models

type Usuario struct {
	ID                        int    `gorm:"column:id;primaryKey;autoIncrement"`
	Id_user_identidad_digital int    `gorm:"column:id_user_identidad_digital"`
	Nombres                   string `gorm:"column:nombres" json:"nombre"`
	Apellidos                 string `gorm:"column:apellidos" json:"apellido"`
	Correo                    string `gorm:"column:correo" json:"correo"`
	Passwd                    string `gorm:"column:password" json:"passwd"`
	Documento                 string `gorm:"column:documento"`
}

// TableName sobrescribe el nombre de la tabla usado por Usuario para `mnt_usuario`.
func (Usuario) TableName() string {
	return "mnt_usuario"
}
