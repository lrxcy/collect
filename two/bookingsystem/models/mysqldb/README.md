# Many To Many
Many to Many adds a join table between two models.

For example, if your application includes users and languages,

and a user can speak many languages, and many users can speak a specfied language.

```go
// User has and belongs to many languages, use `user_languages` as join table
type User struct {
  gorm.Model
  Languages         []*Language `gorm:"many2many:user_languages;"`
}

type Language struct {
  gorm.Model
  Name string
  Users         	  []*User     `gorm:"many2many:user_languages;"`
}
```

create some record for the many-to-many tables
```go
var users []User
language := Language{}

db.First(&language, "id = ?", 111)

db.Model(&language).Related(&users,  "Users")

//// SELECT * FROM "users" INNER JOIN "user_languages" ON "user_languages"."user_id" = "users"."id" WHERE  ("user_languages"."language_id" IN ('111'))
```


# refer:
- http://gorm.io/docs/many_to_many.html
- https://github.com/jinzhu/gorm/issues/754


# 如何使用gorm轉換為sql syntax
- https://www.wancat.cc/2019/07/26/orm/

# sql create many-to-many
- https://stackoverflow.com/questions/7296846/how-to-implement-one-to-one-one-to-many-and-many-to-many-relationships-while-de

# insert a record into an existing has many table with gorm
- https://stackoverflow.com/questions/30969017/how-to-insert-a-record-into-an-existing-has-many-table-with-gorm