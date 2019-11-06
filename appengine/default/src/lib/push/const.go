package push

// Platform ... 送信元のプラットフォーム
type Platform string

const (
	// PlatformIOS ... iOS
	PlatformIOS Platform = "ios"
	// PlatformAndroid ... Android
	PlatformAndroid Platform = "android"
	// PlatformWeb ... Web
	PlatformWeb Platform = "web"
)
