// internal/domain/models/diaries.go
package models

type Diaries struct {
    diaries []*Diary
}

func NewDiaries() *Diaries {
    return &Diaries{
        diaries: make([]*Diary, 0),
    }
}

func (ds *Diaries) Add(diary *Diary) {
    ds.diaries = append(ds.diaries, diary)
}

func (ds *Diaries) FindByUsername(username string) *Diary {
    for _, diary := range ds.diaries {
        if diary.Username == username {
            return diary
        }
    }
    return nil
}

func (ds *Diaries) Delete(username, password string) bool {
    for i, diary := range ds.diaries {
        if diary.Username == username && diary.Password == password {
            ds.diaries = append(ds.diaries[:i], ds.diaries[i+1:]...)
            return true
        }
    }
    return false
}