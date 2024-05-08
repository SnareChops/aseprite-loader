package internal

type ExternalFiles struct {
	Entries  []ExternalFile
	UserData UserData
}

type ExternalFile struct {
	EntryID uint32
	Type    byte
	Name    string
}
