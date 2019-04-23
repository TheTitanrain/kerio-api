package client

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

// ABExtension Extension of Apple Address Book
type ABExtension struct {
	GroupID string `json:"groupId"`
	Label   string `json:"label"`
}

type AccountSyncKey struct {
	Guid      string `json:"guid"`
	Watermark string `json:"watermark"`
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
