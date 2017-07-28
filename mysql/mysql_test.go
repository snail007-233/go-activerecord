package mysql

import (
	"strings"
	"testing"
)

func ar() *ActiveRecord {
	ar := new(ActiveRecord)
	ar.Reset()
	return ar
}
func TestFrom(t *testing.T) {
	want := "SELECT * \nFROM `test`"
	got := strings.TrimSpace(ar().From("test").SQL())
	if want != got {
		t.Errorf("TestFrom , except:%s , got:%s", want, got)
	}
}
func TestFromAs(t *testing.T) {
	want := "SELECT * \nFROM `test` AS `asname`"
	got := strings.TrimSpace(ar().FromAs("test", "asname").SQL())
	if want != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}

func TestSelect(t *testing.T) {
	want := "SELECT `a`,`b` \nFROM `test`"
	got := strings.TrimSpace(ar().From("test").Select("a,b").SQL())
	if want != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}
func TestJoin(t *testing.T) {
	want := "SELECT `u`.`a`,`test`.`b` \nFROM `test` LEFT JOIN `user` AS `u` ON `u`.`a`=`test`.`a`"
	got := strings.TrimSpace(ar().From("test").Select("u.a,test.b").Join("user", "u", "u.a=test.a", "LEFT").SQL())
	if want != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}
func TestWhere(t *testing.T) {
	_ar := ar()
	want := "SELECT * \nFROM `test` \nWHERE `addr` = ? AND `name` = ?"
	want1 := "SELECT * \nFROM `test` \nWHERE `name` = ? AND `addr` = ?"
	got := strings.TrimSpace(_ar.From("test").Where(map[string]interface{}{
		"name": "kitty",
		"addr": "hk",
	}).SQL())
	if want != got && want1 != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}
func TestGroupBy(t *testing.T) {
	want := "SELECT * \nFROM `test`  \nGROUP BY `name`,`uid`"
	got := strings.TrimSpace(ar().From("test").GroupBy("name,uid").SQL())
	if want != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}
func TestHaving(t *testing.T) {
	want := "SELECT * \nFROM `test`  \nGROUP BY `name`,`uid` \nHAVING count(uid)>3"
	got := strings.TrimSpace(ar().From("test").GroupBy("name,uid").Having("count(uid)>3").SQL())
	if want != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}

func TestOrderBy(t *testing.T) {
	want := "SELECT * \nFROM `test`    \nORDER BY `id` DESC,`name` ASC"
	got := strings.TrimSpace(ar().From("test").OrderBy("id", "desc").OrderBy("name", "asc").SQL())
	if want != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}
func TestLimit(t *testing.T) {
	want := "SELECT * \nFROM `test`     \nLIMIT 0,3"
	got := strings.TrimSpace(ar().From("test").Limit(0, 3).SQL())
	if want != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}
func TestInsert(t *testing.T) {
	_ar := ar()
	want := "INSERT INTO  `test` (`name`,`gid`,`addr`,`is_delete`) \nVALUES (?,?,?,?)"
	got := strings.TrimSpace(_ar.Insert("test", map[string]interface{}{
		"name":      "admin",
		"gid":       33,
		"addr":      nil,
		"is_delete": false,
	}).Limit(0, 3).SQL())
	//fmt.Println(_ar.Values())
	if want != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}
func TestReplace(t *testing.T) {
	_ar := ar()
	want := "REPLACE INTO  `test` (`name`,`gid`,`addr`,`is_delete`) \nVALUES (?,?,?,?)"
	got := strings.TrimSpace(_ar.Replace("test", map[string]interface{}{
		"name":      "admin",
		"gid":       33,
		"addr":      nil,
		"is_delete": false,
	}).Limit(0, 3).SQL())
	//fmt.Println(_ar.Values())
	if want != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}

func TestInsertBatch(t *testing.T) {
	_ar := ar()
	want := "INSERT INTO  `test` (`name`) \nVALUES (?),(?)"
	got := strings.TrimSpace(_ar.InsertBatch("test", []map[string]interface{}{
		map[string]interface{}{
			"name": "admin11",
		},
		map[string]interface{}{
			"name": "admin",
		},
	}).SQL())

	if want != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}
func TestReplaceBatch(t *testing.T) {
	_ar := ar()
	want := "REPLACE INTO  `test` (`name`) \nVALUES (?),(?)"
	got := strings.TrimSpace(_ar.ReplaceBatch("test", []map[string]interface{}{
		map[string]interface{}{
			"name": "admin11",
		},
		map[string]interface{}{
			"name": "admin",
		},
	}).SQL())

	if want != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}
func TestDelete(t *testing.T) {
	want := "DELETE FROM  `test`"
	got := strings.TrimSpace(ar().Delete("test", nil).SQL())
	if want != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}
func TestUpdate(t *testing.T) {
	_ar := ar()
	want := "UPDATE  `test` \nSET `addr` = NULL"
	got := strings.TrimSpace(_ar.Update("test", map[string]interface{}{
		"addr": nil,
	}, nil).SQL())
	//fmt.Println(_ar.Values())
	if want != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}

func TestUpdateBatch(t *testing.T) {
	_ar := ar()
	want := "UPDATE  `test` \nSET `name` = CASE \nWHEN `gid` = ? THEN ? \nWHEN `gid` = ? THEN ? \nELSE `name` END \nWHERE gid IN (?,?)"
	got := strings.TrimSpace(_ar.UpdateBatch("test", []map[string]interface{}{
		map[string]interface{}{
			"name": "admin11",
			"gid":  22,
		},
		map[string]interface{}{
			"name": "admin",
			"gid":  33,
		},
	}, "gid").SQL())
	//fmt.Println(_ar.Values())
	if want != got {
		t.Errorf("\n==> Except : \n%s\n==> Got : \n%s", want, got)
	}
}
