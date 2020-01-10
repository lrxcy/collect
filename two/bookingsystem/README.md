# 搶票系統
### 業務流程 & 需求
1. 帳號持有者從前臺，夾帶驗證token進入請求
2. 請求需要購買的套票種類
3. 購買成功寫入DB

### 架構流程
```
nginx -> service(s) -> redis(setnx) -> mysql
```
request(token) -> parse(token) -> purchase(redis_setnx_expire) -> purchase(billing_interface) -> finish(billing_invoice) -> checkout(redis_setnx)-> mysql

### notes
1. token驗證
2. 票數限制SETNX ticket_num "value"
3. 交易成功後刪除key，DEL ticket_num


# refer
- https://github.com/gomodule/redigo
- https://redis.io/commands/setnx

# create a dummy mysql database
```sql
create user 'jim' IDENTIFIED by 'password';

create database `demodatabase`;
grant all privileges on demodatabase.* to 'jim';
```