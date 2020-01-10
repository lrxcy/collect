# Memo Server
This is an memo server used to add my notes ...

# feature

# execute with cli
### post data
```
curl --request POST \
  --url http://127.0.0.1:8080/v1 \
  --header 'content-type: application/json' \
  --data '{
	"Title":"test1",
	"Description": "test description",
	"Category": 0
}'
```

### get data
```
curl --request GET \
  --url 'http://127.0.0.1:8080/v1?page=1&limit=20'
```

### update data
```
curl --request PUT \
  --url http://127.0.0.1:8080/v1 \
  --header 'content-type: application/json' \
  --data '{
	"data":
	[
		{
			"ID" : 1,
			"Title": "testTitle10",
			"Description": "testDescription123",
			"Category":1	
		},
		{
			"ID" : 2,
			"Title": "testTitle9",
			"Description": "testDescription456",
			"Category":1	
		}
	]
}'
```

### delete data
```
curl --request DELETE \
  --url http://127.0.0.1:8080/v1 \
  --header 'content-type: application/json' \
  --data '{
	"data":
	[
		{"ID": 1},
		{"ID": 2}
	]
}'
```


# refer:

### paginator reference
- https://github.com/biezhi/gorm-paginator

### way to check gin request body
- https://github.com/gin-gonic/gin/issues/1295
- https://blog.csdn.net/Manrener/article/details/52182713
- https://stackoverflow.com/questions/32008680/gin-go-lang-how-to-use-context-request-body-and-retain-it

### batch update:
- http://gorm.io/docs/update.html
