package stringlist

import (
	"reflect"
	"testing"
)

var testRows = []string{
	"asd adslk d asdlkas dlkaslk dl aks",
	"asda sdasj da js a s jas jk kja sdk",
	"ыва ДЛ о  ывло дДЛО  ЛДО дл ДЛ ",
	"sadjf hsdkjfh kjhhalsd fkjhas l alkjsdhf aj ы двлоадф213о 409з - - 0ыфваыв09ал ыщаз9ващылв",
	"asd adslk d asdlkas dlkaslk dl aks",
	"asda sdasj da js a s jas jk kja sdk",
	"ыва ДЛ о  ывло дДЛО  ЛДО дл ДЛ ",
	"sadjf hsdkjfh kjhhalsd fkjhas l alkjsdhf aj ы двлоадф213о 409з - - 0ыфваыв09ал ыщаз9ващылв",
	"asd adslk d asdlkas dlkaslk dl aks",
	"asda sdasj da js a s jas jk kja sdk",
	"ыва ДЛ о  ывло дДЛО  ЛДО дл ДЛ ",
	"sadjf hsdkjfh kjhhalsd fkjhas l alkjsdhf aj ы двлоадф213о 409з - - 0ыфваыв09ал ыщаз9ващылв",
	"asd adslk d asdlkas dlkaslk dl aks",
	"asda sdasj da js a s jas jk kja sdk",
	"ыва ДЛ о  ывло дДЛО  ЛДО дл ДЛ ",
	"sadjf hsdkjfh kjhhalsd fkjhas l alkjsdhf aj ы двлоадф213о 409з - - 0ыфваыв09ал ыщаз9ващылв",
	"asd adslk d asdlkas dlkaslk dl aks",
	"asda sdasj da js a s jas jk kja sdk",
	"ыва ДЛ о  ывло дДЛО  ЛДО дл ДЛ ",
	"sadjf hsdkjfh kjhhalsd fkjhas l alkjsdhf aj ы двлоадф213о 409з - - 0ыфваыв09ал ыщаз9ващылв",
	"asd adslk d asdlkas dlkaslk dl aks",
	"asda sdasj da js a s jas jk kja sdk",
	"ыва ДЛ о  ывло дДЛО  ЛДО дл ДЛ ",
	"sadjf hsdkjfh kjhhalsd fkjhas l alkjsdhf aj ы двлоадф213о 409з - - 0ыфваыв09ал ыщаз9ващылв",
}

// 24
func Test_stringList_Add(t *testing.T) {
	t.Run("cap < len of inserting", func(t *testing.T) {
		sl := New(0)
		for i := 0; i < len(testRows); i++ {
			sl.Add(testRows[i])
		}
		if sl.Len() != len(testRows) {
			t.Errorf("Add(): want len %d, got %d", len(testRows), sl.Len())
		}
		if !reflect.DeepEqual(sl.Rows(), testRows) {
			t.Errorf("Add(): want rows %v, got %v", testRows, sl.Rows())
		}
		for i := 0; i < len(testRows); i++ {
			if val, err := sl.RowAt(i); val != testRows[i] {
				t.Errorf("Add(): wrong result at pos %d, want %s, got %s", i, testRows[i], val)
				if err != nil {
					t.Errorf("Add(): %v", err)
				}
			}
		}
		if _, err := sl.RowAt(len(testRows) + 1); err == nil {
			t.Errorf("Add(): want index out of bounds, got no error (to upper side)")
		}
		if _, err := sl.RowAt(-1); err == nil {
			t.Errorf("Add(): want index out of bounds, got no error (to lower side)")
		}

		if cap(sl.rws) != 30 {
			t.Errorf("Add(): want inner slice cap: %d, got %d", 25, cap(sl.rws))
		}

		sl = New(60)
		for i := 0; i < len(testRows); i++ {
			sl.Add(testRows[i])
		}
		if sl.Len() != len(testRows) {
			t.Errorf("Add(): want len %d, got %d", len(testRows), sl.Len())
		}
		if !reflect.DeepEqual(sl.Rows(), testRows) {
			t.Errorf("Add(): want rows %v, got %v", testRows, sl.Rows())
		}
		for i := 0; i < len(testRows); i++ {
			if val, err := sl.RowAt(i); val != testRows[i] {
				t.Errorf("Add(): wrong result at pos %d, want %s, got %s", i, testRows[i], val)
				if err != nil {
					t.Errorf("Add(): %v", err)
				}
			}
		}
		if _, err := sl.RowAt(len(testRows) + 1); err == nil {
			t.Errorf("Add(): want index out of bounds, got no error (to upper side)")
		}
		if _, err := sl.RowAt(-1); err == nil {
			t.Errorf("Add(): want index out of bounds, got no error (to lower side)")
		}

		if cap(sl.rws) != 60 {
			t.Errorf("Add(): want inner slice cap: %d, got %d", 25, cap(sl.rws))
		}

		rows1 := sl.Rows()
		rows1[0] = "qqqqqqqqq"
		if val, _ := sl.RowAt(0); rows1[0] == val {
			t.Errorf("sl.Rows() returns the same lice")
		}
	})
}

func Test_stringList_AddRows(t *testing.T) {
	t.Run("cap < len of inserting", func(t *testing.T) {
		sl := New(0)
		sl.AddRows(testRows)
		if sl.Len() != len(testRows) {
			t.Errorf("Add(): want len %d, got %d", len(testRows), sl.Len())
		}
		if !reflect.DeepEqual(sl.Rows(), testRows) {
			t.Errorf("Add(): want rows %v, got %v", testRows, sl.Rows())
		}
		for i := 0; i < len(testRows); i++ {
			if val, err := sl.RowAt(i); val != testRows[i] {
				t.Errorf("Add(): wrong result at pos %d, want %s, got %s", i, testRows[i], val)
				if err != nil {
					t.Errorf("Add(): %v", err)
				}
			}
		}
		if _, err := sl.RowAt(len(testRows) + 1); err == nil {
			t.Errorf("Add(): want index out of bounds, got no error (to upper side)")
		}
		if _, err := sl.RowAt(-1); err == nil {
			t.Errorf("Add(): want index out of bounds, got no error (to lower side)")
		}

		if cap(sl.rws) != 34 {
			t.Errorf("Add(): want inner slice cap: %d, got %d", 34, cap(sl.rws))
		}

		sl = New(60)
		sl.AddRows(testRows)
		if sl.Len() != len(testRows) {
			t.Errorf("Add(): want len %d, got %d", len(testRows), sl.Len())
		}
		if !reflect.DeepEqual(sl.Rows(), testRows) {
			t.Errorf("Add(): want rows %v, got %v", testRows, sl.Rows())
		}
		for i := 0; i < len(testRows); i++ {
			if val, err := sl.RowAt(i); val != testRows[i] {
				t.Errorf("Add(): wrong result at pos %d, want %s, got %s", i, testRows[i], val)
				if err != nil {
					t.Errorf("Add(): %v", err)
				}
			}
		}
		if _, err := sl.RowAt(len(testRows) + 1); err == nil {
			t.Errorf("Add(): want index out of bounds, got no error (to upper side)")
		}
		if _, err := sl.RowAt(-1); err == nil {
			t.Errorf("Add(): want index out of bounds, got no error (to lower side)")
		}

		if cap(sl.rws) != 60 {
			t.Errorf("Add(): want inner slice cap: %d, got %d", 25, cap(sl.rws))
		}
	})
}

func Test_stringList_AddFrom(t *testing.T) {
	type args struct {
		list Strings
	}
	tests := []struct {
		name string
		sl   *StringList
		args args
	}{
		//  TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sl.AddFrom(tt.args.list)
		})
	}
}

func Test_stringList_At(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		sl      *StringList
		args    args
		want    string
		wantErr bool
	}{
		//  TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.sl.RowAt(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("stringList.At() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("stringList.At() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringList_Rows(t *testing.T) {
	tests := []struct {
		name string
		sl   *StringList
		want []string
	}{
		//  TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sl.Rows(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringList.Rows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringList_Text(t *testing.T) {
	tests := []struct {
		name string
		sl   *StringList
		want string
	}{
		//  TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sl.Text(); got != tt.want {
				t.Errorf("stringList.Text() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringList_Join(t *testing.T) {
	type args struct {
		sep string
	}
	tests := []struct {
		name string
		sl   *StringList
		args args
		want string
	}{
		//  TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sl.Join(tt.args.sep); got != tt.want {
				t.Errorf("stringList.Join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringList_SetLineSep(t *testing.T) {
	type args struct {
		sep string
	}
	tests := []struct {
		name string
		sl   *StringList
		args args
	}{
		//  TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sl.SetLineSep(tt.args.sep)
		})
	}
}
