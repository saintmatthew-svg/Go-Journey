package main

type Diaries struct {
	diaries []*Diary
}

func (d *Diaries) Add(diary *Diary) {
	d.diaries = append(d.diaries, diary)
}

func (d *Diaries) FindByUsername(username string) *Diary {
	for _, diary := range d.diaries {
		if diary.username == username {
			return diary
		}
	}
	return nil
}

func (d *Diaries) Delete(username string) {
	for i, diary := range d.diaries {
		if diary.username == username {
			d.diaries = append(d.diaries[:i], d.diaries[i+1:]...)
			return
		}
	}
}