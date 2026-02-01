package siq

// Content types
const (
	ContentTypeText  = "text"
	ContentTypeImage = "image"
	ContentTypeAudio = "audio"
	ContentTypeVideo = "video"
	ContentTypeHTML  = "html"
)

// Content placements
const (
	PlacementScreen     = "screen"
	PlacementReplic     = "replic"
	PlacementBackground = "background"
)

// Round types
const (
	RoundTypeStandard = "standart" // Note: typo preserved for compatibility
	RoundTypeFinal    = "final"
)

// Question types
const (
	QuestionTypeDefault           = "default"
	QuestionTypeStake             = "stake"
	QuestionTypeSecret            = "secret"
	QuestionTypeSecretPublicPrice = "secretPublicPrice"
	QuestionTypeSecretNoQuestion  = "secretNoQuestion"
	QuestionTypeNoRisk            = "noRisk"
	QuestionTypeForAll            = "forAll"
)

// Step parameter types
const (
	StepParamTypeSimple    = "simple"
	StepParamTypeContent   = "content"
	StepParamTypeGroup     = "group"
	StepParamTypeNumberSet = "numberSet"
)

// Collection names for media storage
const (
	CollectionImages = "Images"
	CollectionAudio  = "Audio"
	CollectionVideo  = "Video"
	CollectionHTML   = "Html"
)
// Question parameter names
const (
	ParamQuestion      = "question"
	ParamAnswer        = "answer"
	ParamAnswerType    = "answerType"
	ParamAnswerOptions = "answerOptions"
)

// Package version
const (
	PackageVersion = 5.0
)
