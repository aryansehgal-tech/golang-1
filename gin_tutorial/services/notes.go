package services

type NotesServices struct {

}

type Note struct {
	Id int
	Name string
}

func (n *NotesServices) GetNotesServices() []Note {
	data:=[]Note{
		{Id:1,Name:"Note 1"},
		{Id:2,Name:"Note 2"},
	}
	return data
}

func (n *NotesServices) CreateNotesServices() []Note {
	data:=[]Note{
		{Id:3, Name:"Note 3"},
	}
	return data
}