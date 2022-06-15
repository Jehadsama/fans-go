package models

type OaPGfUser struct {
	BusiDate          string `gorm:"type:varchar(255);" json:"busi_date"`
	Userid            string `gorm:"type:varchar(255);" json:"user_id"`
	Loginid           string `gorm:"type:varchar(255);" json:"loginid"`
	Erpid             string `gorm:"type:varchar(255);" json:"erpid"`
	Username          string `gorm:"type:varchar(255);" json:"username"`
	Initialcharacter  string `gorm:"type:varchar(255);" json:"initialcharacter"`
	Dpid              string `gorm:"type:varchar(255);" json:"Dpid"`
	Phone             string `gorm:"type:varchar(255);" json:"phone"`
	Directphone       string `gorm:"type:varchar(255);" json:"directphone"`
	Mobile            string `gorm:"type:varchar(255);" json:"mobile"`
	Mailbox           string `gorm:"type:varchar(255);" json:"mailbox"`
	Sex               string `gorm:"type:varchar(255);" json:"sex"`
	Serialnumber      string `gorm:"type:varchar(255);" json:"serialnumber"`
	Qqormsn           string `gorm:"type:varchar(255);" json:"qqormsn"`
	Workingplace      string `gorm:"type:varchar(255);" json:"workingplace"`
	Post              string `gorm:"type:varchar(255);" json:"post"`
	Rank              string `gorm:"type:varchar(255);" json:"rank"`
	Birthday          string `gorm:"type:varchar(255);" json:"birthday"`
	Employeddate      string `gorm:"type:varchar(255);" json:"employeddate"`
	Dimissiondate     string `gorm:"type:varchar(255);" json:"dimissiondate"`
	Status            string `gorm:"type:varchar(255);" json:"status"`
	Comments          string `gorm:"type:varchar(255);" json:"comments"`
	Photograph        string `gorm:"type:varchar(255);" json:"photograph"`
	Maindepartment    string `gorm:"type:varchar(255);" json:"maindepartment"`
	Superiorleader    string `gorm:"type:varchar(255);" json:"superiorleader"`
	Stafftype         string `gorm:"type:varchar(255);" json:"stafftype"`
	Idiograph         string `gorm:"type:varchar(255);" json:"idiograph"`
	Remarks           string `gorm:"type:varchar(255);" json:"remarks"`
	Oaid              string `gorm:"type:varchar(255);" json:"oaid"`
	Dpdn              string `gorm:"type:varchar(255);" json:"dpdn"`
	Displayname       string `gorm:"type:varchar(255);" json:"displayname"`
	Dn                string `gorm:"type:varchar(255);" json:"dn"`
	Tamuid            string `gorm:"type:varchar(255);" json:"tamuid"`
	Oamailserver      string `gorm:"type:varchar(255);" json:"oamailserver"`
	Oamailfile        string `gorm:"type:varchar(255);" json:"oamailfile"`
	Oacomment         string `gorm:"type:varchar(255);" json:"oacomment"`
	Nationalid        string `gorm:"type:varchar(255);" json:"nationalid"`
	Uniusercode       string `gorm:"type:varchar(255);" json:"uniusercode"`
	Orgcodeofuser     string `gorm:"type:varchar(255);" json:"orgcodeofuser"`
	Closeaccount      string `gorm:"type:varchar(255);" json:"closeaccount"`
	Usercreatedate    string `gorm:"type:varchar(255);" json:"usercreatedate"`
	Usermodifydate    string `gorm:"type:varchar(255);" json:"usermodifydate"`
	Userdeletedate    string `gorm:"type:varchar(255);" json:"userdeletedate"`
	Positioncatcode   string `gorm:"type:varchar(255);" json:"positioncatcode"`
	Positioncat       string `gorm:"type:varchar(255);" json:"positioncat"`
	Position          string `gorm:"type:varchar(255);" json:"position"`
	Orgidofuser       string `gorm:"type:varchar(255);" json:"orgidofuser"`
	Unideptcodeofuser string `gorm:"type:varchar(255);" json:"unideptcodeofuser"`
	DataTime          string `gorm:"type:varchar(255);" json:"data_time"`
}

// TableName 会将 User 的表名重写为 `profiles`
func (OaPGfUser) TableName() string {
	return "oa_p_gf_user"
}

func (user *OaPGfUser) UsernameByOAID(oaID string) string {
	result := Model(&user).Select("username").Where("loginid = ?", oaID).First(&user)
	if result.RowsAffected == 1 {
		return user.Username
	}
	return ""
}

func OaPGfUserModel() *OaPGfUser {
	return &OaPGfUser{}
}
