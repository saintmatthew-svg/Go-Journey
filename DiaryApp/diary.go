package main

type Diary struct {
	username string
	password string
	isLocked bool
	entries  []*Entry
}

func NewDiary(username, password string) *Diary {
	return &Diary{
		username: username,
		password: password,
		isLocked: true,
		entries:  []*Entry{},
	}
}


func (d *Diary) UnlockDiary(password string) {
	if d.password == password {
		d.isLocked = false
	}
}


func (d *Diary) LockDiary() {
	d.isLocked = true
}


func (d *Diary) IsLocked() bool {
	return d.isLocked
}


func (d *Diary) CreateEntry(title, body string) {
	if !d.isLocked {
		id := len(d.entries) + 1 
		entry := NewEntry(id, title, body)
		d.entries = append(d.entries, entry)
	}
}


func (d *Diary) DeleteEntry(id int) {
	if !d.isLocked {
		for i, entry := range d.entries {
			if entry.ID == id {
				d.entries = append(d.entries[:i], d.entries[i+1:]...)
				return
			}
		}
	}
}



func (d *Diary) FindEntryByID(id int) *Entry {
	if !d.isLocked {
		for _, entry := range d.entries {
			if entry.ID == id {
				return entry
			}
		}
	}
	return nil
}


func (d *Diary) UpdateEntry(id int, title, body string) {
	if !d.isLocked {
		for _, entry := range d.entries {
			if entry.ID == id {
				entry.Title = title
				entry.Body = body
				return
			}
		}
	}
}