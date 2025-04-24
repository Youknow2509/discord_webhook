package model

// UserObject represents a Discord user
type UserObject struct {
    ID                  string                `json:"id"`
    Username            string                `json:"username"`
    Discriminator       string                `json:"discriminator"`
    GlobalName          *string               `json:"global_name,omitempty"`
    Avatar              *string               `json:"avatar,omitempty"`
    Bot                 bool                  `json:"bot,omitempty"`
    System              bool                  `json:"system,omitempty"`
    MFAEnabled          bool                  `json:"mfa_enabled,omitempty"`
    Banner              *string               `json:"banner,omitempty"`
    AccentColor         *int                  `json:"accent_color,omitempty"`
    Locale              string                `json:"locale,omitempty"`
    Verified            bool                  `json:"verified,omitempty"`
    Email               *string               `json:"email,omitempty"`
    Flags               int                   `json:"flags,omitempty"`
    PremiumType         int                   `json:"premium_type,omitempty"`
    PublicFlags         int                   `json:"public_flags,omitempty"`
    AvatarDecorationData *AvatarDecorationData `json:"avatar_decoration_data,omitempty"`
}

// AvatarDecorationData represents decoration data for a user's avatar
type AvatarDecorationData struct {
    Asset      string `json:"asset,omitempty"`
    SkuID      string `json:"sku_id,omitempty"`
}

// Premium types
type PremiumType int

const (
    NoPremium PremiumType = iota
    NitroClassic
    Nitro
    NitroBasic
)

// User flags
type UserFlag int

const (
    Staff UserFlag = 1 << iota
    Partner
    HypeSquad
    BugHunterLevel1
    // ...
)