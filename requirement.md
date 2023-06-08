# Project Requirement

## Talbe
buatlah sebuah table dengan nama tasks dengan kolom
- ID serial primary key
- title varchar
- description varchar
- created_at date
- updated_at date
- deleted_at date

<br>

## Endpoint
### <b>GET /tasks</b>
Menampilkan semua tasks yang tersimpan pada database
<br><b>Response</b>
```json
[
  {
    "ID": int,
    "title": string,
    "description": string,
  },
  ...
]
```
<br>

### <b>POST /tasks</b>
Menambahkan task ke dalam database
<br><b>Request</b>
```json
{
  "title": string,
  "description": string
}
```
<b>Response</b>
```json
{
  "message": "success create task"
}
```

<br>

### <b>DELETE /tasks/:id</b>
Menghapus task yang ada pada database sesuai dengan id nya
<br><b>Response</b>
```json
{
  "message": "success delete task"
}
```