package design

import . "goa.design/goa/v3/dsl"

var StudentType = Type("StudentType", func() {
	Description("One student")
	Attributes(func() {
		Attribute("id", Int64, "学生を一意に表す ID", func() {
			Example(1)
		})
		Attribute("name", String, "学生の氏名", func() {
			Example("鈴木太郎")
		})
		Attribute("ruby", String, "学生の氏名のフリガナ", func() {
			Example("スズキタロウ")
		})
		Attribute("student_number", Int, "学生の学籍番号", func() {
			Example(12345)
		})
		Attribute("date_of_birth", FormatDateTime, "学生の生年月日", func() {
			Example("2022-04-01T13:30:00+09:00")
		})
		Attribute("address", String, "学生の住所", func() {
			Example("名古屋市中区三の丸三丁目1番2号")
		})
		Attribute("expiration_date", FormatDateTime, "学生証の有効期間", func() {
			Example("2027-03-31T00:00:00+09:00")
		})
		Required("id", "name", "ruby", "student_number", "date_of_birth", "address", "expiration_date")
		View("default", func() {
			Attribute("id")
			Attribute("name")
			Attribute("ruby")
			Attribute("student_number")
			Attribute("date_of_birth")
			Attribute("address")
			Attribute("expiration_date")
		})
	})
})

var StudentsType = Type("StudentsType", func() {
	Description("All students")
	Attribute("students", ArrayOf(StudentType), func() {
		ArrayOf(StudentType)
	})
	Required("students")
	View("default", func() {
		Attribute("students")
	})
})
