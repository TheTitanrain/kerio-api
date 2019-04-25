package client

type (
	Watermark    int
	Kld          int
	PriorityType string
)

// LocalizableMessage Message can contain replacement marks: { "User %1 cannot be deleted.", ["jsmith"], 1 }
type LocalizableMessage struct {
	Message              string   `json:"message"`
	PasitionalParameters []string `json:"pasitionalParameters"`
	Plurality            int      `json:"plurality"`
}

// ApiException -
type ApiException struct {
	Message           string                       `json:"message"`
	Code              int                          `json:"code"`
	MessageParameters LocalizableMessageParameters `json:"messageParameters"`
}

type LocalizableMessageParameters struct {
	PositionalParameters []string `json:"positionalParameters"`
	Plurality            int      `json:"plurality"`
}

// AddResult - Result of the add operation
type AddResult struct {
	Id           int                `json:"id"`
	Success      bool               `json:"success"`
	ErrorMessage LocalizableMessage `json:"errorMessage"`
}

// CreateResult - Details about a particular item created.
type CreateResult struct {
	InputIndex int       `json:"inputIndex"`
	Id         int       `json:"id"`
	Watermark  Watermark `json:"watermark"`
}

// SetResult - Details about a particular item created.
type SetResult struct {
	InputIndex int       `json:"inputIndex"`
	Id         Kld       `json:"id"`
	Watermark  Watermark `json:"watermark"`
}

// Image - [READ ONLY] Id of uploaded image.
type Image struct {
	Url string `json:"url"`
	Id  string `json:"id"`
}

// Attachment
type Attachment struct {
	Id          string `json:"id"`
	Url         string `json:"url"`
	Name        string `json:"name"`
	ContentType string `json:"contentType"`
	ContentId   string `json:"contentId"`
	Size        int    `json:"size"`
}

type Attendee struct {
	DisplayName string `json:"displayName"`
	Role        string `json:"role"`
	IsNotified  bool   `json:"isNotified"`
	PartStatus  string `json:"partStatus"`
}

type OperatorExtension struct {
	ExtensionId  string `json:"extensionId"`
	TelNum       string `json:"telNum"`
	Description  string `json:"description"`
	IsRegistered bool   `json:"isRegistered"`
}

// ApiApplication - Describes client (third-party) application or script which uses the Administration API
type ApiApplication struct {
	Name    string `json:"name"`
	Vendor  string `json:"vendor"`
	Version string `json:"version"`
}

type Error struct {
	InputIndex        int                          `json:"inputIndex"`
	Code              int                          `json:"code"`
	Message           string                       `json:"message"`
	MessageParameters LocalizableMessageParameters `json:"messageParameters"`
}

// NamedValue pair Note: all fields must be assigned if used in set methods
type NamedValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type DisplayablePart struct {
	ContentType string `json:"contentType"`
	Content     string `json:"content"`
}

type Email struct {
	Address   string `json:"address"`
	Name      string `json:"name"`
	ContactId string `json:"contactId"`
}

type Mail struct {
	From             Email             `json:"from"`
	Sender           Email             `json:"sender"`
	DisplayableParts []DisplayablePart `json:"displayableParts"`
	To               []Email           `json:"to"`
	Cc               []Email           `json:"cc"`
	Bcc              []Email           `json:"bcc"`
	ReplyTo          []Email           `json:"replyTo"`
	NotificationTo   Email             `json:"notificationTo"`
	Subject          string            `json:"subject"`
	Priority         PriorityType      `json:"priority"`
	Size             int               `json:"size"`
	Encrypt          bool              `json:"encrypt"`
	IsAnswered       bool              `json:"isAnswered"`
	IsDraft          bool              `json:"isDraft"`
	IsFlagged        bool              `json:"isFlagged"`
	IsForwarded      bool              `json:"isForwarded"`
	IsJunk           bool              `json:"isJunk"`
	IsMDNSent        bool              `json:"isMdnSent"`
	IsReadOnly       bool              `json:"isReadOnly"`
	IsSeen           bool              `json:"isSeen"`
	RequestDSN       bool              `json:"requestDsn"`
	Send             bool              `json:"send"`
	ShowExternal     bool              `json:"showExternal"`
	Sign             bool              `json:"sign"`
}

type OutIsEligible struct {
	IsEligible bool `json:"isEligible"`
}

// ABExtension Extension of Apple Address Book
type ABExtension struct {
	GroupID string `json:"groupId"`
	Label   string `json:"label"`
}

type AccountSyncKey struct {
	Guid      string    `json:"guid"`
	Watermark Watermark `json:"watermark"`
}

// UserInfo - Details of the logged user into the webmail
type UserInfo struct {
	Id               int      `json:"id"`
	LoginName        string   `json:"loginName"`
	FullName         string   `json:"fullName"`
	Emals            []string `json:"emals"`
	PreferredAddress string   `json:"preferredAddress"`
	ReplyToAddress   string   `json:"replyToAddress"`
}

// OutUserInfo - Details of the logged user into the webmail (for returned structures)
type OutUserInfo struct {
	UserDetails UserInfo `json:"userDetails"`
}
