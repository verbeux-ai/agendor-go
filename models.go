package agendor

type Deal struct {
	Id               int             `json:"id,omitempty"`
	Title            string          `json:"title,omitempty"`
	AccountId        int             `json:"accountId,omitempty"`
	StartTime        any             `json:"startTime,omitempty"`
	EndTime          any             `json:"endTime,omitempty"`
	WonAt            any             `json:"wonAt,omitempty"`
	LostAt           any             `json:"lostAt,omitempty"`
	Organization     Organization    `json:"organization,omitempty"`
	Person           People          `json:"person,omitempty"`
	Author           Author          `json:"author,omitempty"`
	DealStage        DealStage       `json:"dealStage,omitempty"`
	DealStatus       DealStatus      `json:"dealStatus,omitempty"`
	LossReason       LossReason      `json:"lossReason,omitempty"`
	Owner            Owner           `json:"owner,omitempty"`
	Description      string          `json:"description,omitempty"`
	CreatedAt        any             `json:"createdAt,omitempty"`
	UpdatedAt        any             `json:"updatedAt,omitempty"`
	Ranking          int             `json:"ranking,omitempty"`
	Value            float32         `json:"value,omitempty"`
	ProductsEntities []ProductEntity `json:"products_entities,omitempty"`
	Products         []any           `json:"products,omitempty"`
	AllowedUsers     []User          `json:"allowedUsers,omitempty"`
	Email            string          `json:"_email,omitempty"`
	CustomFields     map[string]any  `json:"customFields,omitempty"`
	WebUrl           string          `json:"_webUrl,omitempty"`
}

type Discount float64

type ProductEntity struct {
	Id         int         `json:"id,omitempty"`
	CreatedAt  interface{} `json:"createdAt,omitempty"`
	UnitValue  float32     `json:"unitValue,omitempty"`
	Quantity   int         `json:"quantity,omitempty"`
	Discount   Discount    `json:"discount,omitempty"`
	TotalValue float32     `json:"totalValue,omitempty"`
}

type LossReason struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type DealStage struct {
	Id       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Sequence int    `json:"sequence,omitempty"`
	Funnel   Funnel `json:"funnel,omitempty"`
}

type DealStatus struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Funnel struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type People struct {
	Id           int            `json:"id,omitempty"`
	AccountId    int            `json:"accountId,omitempty"`
	Name         string         `json:"name,omitempty"`
	Email        string         `json:"email,omitempty"`
	Organization Organization   `json:"organization,omitempty"`
	Cpf          string         `json:"cpf,omitempty"`
	Role         string         `json:"role,omitempty"`
	Ranking      int            `json:"ranking,omitempty"`
	Description  string         `json:"description,omitempty"`
	Birthday     string         `json:"birthday,omitempty"`
	Avatar       string         `json:"avatar,omitempty"`
	CreatedAt    interface{}    `json:"createdAt,omitempty"`
	UpdatedAt    interface{}    `json:"updatedAt,omitempty"`
	Products     []any          `json:"products,omitempty"`
	Contact      Contact        `json:"contact,omitempty"`
	Address      Address        `json:"address,omitempty"`
	Category     Category       `json:"category,omitempty"`
	LeadOrigin   LeadOrigin     `json:"leadOrigin,omitempty"`
	Author       Author         `json:"author,omitempty"`
	OwnerUser    Owner          `json:"ownerUser,omitempty"`
	AllowedUsers []User         `json:"allowedUsers,omitempty"`
	CustomFields map[string]any `json:"customFields,omitempty"`
	WebUrl       string         `json:"_webUrl,omitempty"`
}

type User struct {
	AccountId int    `json:"accountId,omitempty"`
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
}

type Owner struct {
	AccountId int    `json:"accountId,omitempty"`
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
}

type Author struct {
	AccountId int    `json:"accountId,omitempty"`
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
}

type LeadOrigin struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Category struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Organization struct {
	Name            string         `json:"name"`
	LegalName       string         `json:"legalName"`
	Cnpj            string         `json:"cnpj"`
	Description     string         `json:"description"`
	Logo            string         `json:"logo"`
	Website         string         `json:"website"`
	Ranking         int            `json:"ranking"`
	OwnerUser       int            `json:"ownerUser"`
	Contact         Contact        `json:"contact"`
	Address         Address        `json:"address"`
	LeadOrigin      int            `json:"leadOrigin"`
	Category        int            `json:"category"`
	Sector          int            `json:"sector"`
	Products        []any          `json:"products"`
	AllowedUsers    []int          `json:"allowedUsers"`
	AllowToAllUsers bool           `json:"allowToAllUsers"`
	CustomFields    map[string]any `json:"customFields"`
}

type Address struct {
	PostalCode     string `json:"postal_code,omitempty"`
	Country        string `json:"country,omitempty"`
	District       string `json:"district,omitempty"`
	State          string `json:"state,omitempty"`
	StreetName     string `json:"street_name,omitempty"`
	StreetNumber   int    `json:"street_number,omitempty"`
	AdditionalInfo string `json:"additional_info,omitempty"`
	City           int    `json:"city,omitempty"`
}

type Contact struct {
	Email     string `json:"email,omitempty"`
	Work      string `json:"work,omitempty"`
	Mobile    string `json:"mobile,omitempty"`
	Fax       string `json:"fax,omitempty"`
	Whatsapp  string `json:"whatsapp,omitempty"`
	Facebook  string `json:"facebook,omitempty"`
	Twitter   string `json:"twitter,omitempty"`
	Instagram string `json:"instagram,omitempty"`
	LinkedIn  string `json:"linked_in,omitempty"`
	Skype     string `json:"skype,omitempty"`
}
