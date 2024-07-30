package models

type ResourceBaseEntity struct {
	Id        int
	Name      string
	ShortName string
	SortOrder int
}

type ResourceDelimiter struct {
	Id        int
	Name      string
	Delimiter string
	Enabled   bool
	SortOrder int
}

type ResourceEnvironment struct {
	ResourceBaseEntity
}

type ResourceFunction struct {
	ResourceBaseEntity
}

type ResourceOrg struct {
	ResourceBaseEntity
}

type ResourceProjAppSvc struct {
	ResourceBaseEntity
}

type ResourceLocation struct {
	ResourceBaseEntity
}

type ResourceUnitDept struct {
	ResourceBaseEntity
}

type ResourceType struct {
	Id                           int
	Resource                     string
	Optional                     string
	Exclude                      string
	Property                     string
	ShortName                    string
	Scope                        string
	LenghtMin                    string
	LenghtMax                    string
	ValidText                    string
	InvalidText                  string
	InvalidCharacters            string
	InvalidCharactersStart       string
	InvalidCharactersEnd         string
	InvalidCharactersConsecutive string
	Regx                         string
	StaticValues                 string
	Enabled                      bool
	ApplyDelimiter               bool
}

type ResourceNameRequest struct {
	ResourceEnvironment string
	ResourceFunction    string
	ResourceInstance    string
	ResourceLocation    string
	ResourceOrg         string
	ResourceProjAppSvc  string
	ResourceType        string
	ResourceUnitDept    string
	CustomComponents    map[string]string
	ResourceId          int
	CreatedBy           string
}

type ResourceNameRequestWithComponents struct {
	ResourceEnvironment ResourceEnvironment
	ResourceFunction    ResourceFunction
	ResourceDelimiter   ResourceDelimiter
	ResourceInstance    string
	ResourceLocation    ResourceLocation
	ResourceOrg         ResourceOrg
	ResourceProjAppSvc  ResourceProjAppSvc
	ResourceType        ResourceType
	ResourceUnitDept    ResourceUnitDept
}

type ResourceNameDetails struct {
	Id               int
	CreatedOn        string
	ResourceName     string
	ResourceTypeName string
	User             string
	Message          string
}

type ResourceNameResponse struct {
	ResourceName        string
	Message             string
	Success             bool
	ResourceNameDetails ResourceNameDetails
}
