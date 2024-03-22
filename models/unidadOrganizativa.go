package models

type UnidadOrganizativa struct {
	ID                int                  `gorm:"column:id;primaryKey;autoIncrement"`
	Codigo            string               `gorm:"column:codigo"`
	Nombre            string               `gorm:"column:nombre"`
	IdEstablecimiento int                  `gorm:"column:id_establecimiento"`
	IdUnidadPadre     int                  `gorm:"column:id_unidad_padre"`
	Activo            bool                 `gorm:"column:activo"`
	NodoPadre         string               `gorm:"column:nodo_padre"`
	Telefono          string               `gorm:"column:telefono"`
	IdMunicipio       int                  `gorm:"column:id_municipio"`
	Direccion         string               `gorm:"column:direccion"`
	IdTipoUnidad      int                  `gorm:"column:id_tipo_unidad"`
	Hijos             []UnidadOrganizativa `gorm:"foreignKey:IdUnidadPadre"`
}

func (UnidadOrganizativa) TableName() string {
	return "mnt_unidad_organizativa"
}
