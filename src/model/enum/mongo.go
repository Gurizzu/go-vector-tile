package enum

type MongoCollection int

const (
	MongoCollection_User MongoCollection = iota
	MongoCollection_Role
	MongoCollection_JobTitle
	MongoCollection_Cards
	MongoCollection_CardLinkType
	MongoCollection_Experiences
)

func (index MongoCollection) String() string {
	return []string{
		"user",
		"role",
		"jobTitle",
		"cards",
		"cardLinkType",
		"experiences",
	}[index]
}
