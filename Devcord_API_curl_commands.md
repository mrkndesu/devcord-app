# 🚀 Devcord APIまとめ（curlコマンド）

## 👤 ユーザー関連API
### 🆕 ユーザー作成（POST /users
```
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "handle": "@Devcord_Japan",
    "name": "デブちゃん",
    "email": "devchan@devcordblog.com",
    "password": "password123",
    "avatar_url": "https://devcordblog.com/avatar.png",
    "description": "はじめまして、私はデブちゃん",
    "birth_date": "2001-01-01",
    "created_year": 2025,
    "created_month": 7
  }'
```

### 👥 全ユーザー取得（GET /users）
```
curl http://localhost:8080/users
```

### 🔍 指定ユーザー取得（GET /users/:userID）
```
curl http://localhost:8080/users/<userID>
```

### ✏️ 指定ユーザー更新（PUT /users/:userID）
```
curl -X PUT http://localhost:8080/users/<userID> \
  -H "Content-Type: application/json" \
  -d '{
    "name": "デブちゃん@プログラマー",
    "description": "私の名前はデブちゃん。自己紹介を更新しました"
  }'
```

### ❌ 指定ユーザー削除（DELETE /users/:userID）
🚨ユーザーを作成すると全投稿を作成される
```
curl -X DELETE http://localhost:8080/users/<userID>
```

## 📝 投稿関連API

### 🆕 投稿作成（POST /users/:userID/posts）
```
curl -X POST http://localhost:8080/users/<userID>/posts \
  -H "Content-Type: application/json" \
  -d '{
    "title": "初めての投稿",
    "content": "みんなはじめまして！"
  }'
```

### 👀 全投稿取得（GET /users/:userID/posts）
```
curl http://localhost:8080/users/<userID>/posts
```

### 🔎 投稿1件取得（GET /users/:userID/posts/:postID）
```
curl http://localhost:8080/users/<userID>/posts/<postID>
```

### ✍️ 投稿更新（PUT /users/:userID/posts/:postID）
```
curl -X PUT http://localhost:8080/users/<userID>/posts/<postID> \
  -H "Content-Type: application/json" \
  -d '{
    "title": "みんな久しぶり！",
    "content": "久しぶりに内容を更新しました"
  }'
```

### 🗑️ 投稿削除（DELETE /users/:userID/posts/:postID）
```
curl -X DELETE http://localhost:8080/users/<userID>/posts/<postID>
```

### 🧹 全投稿削除（DELETE /users/:userID/posts）
```
curl -X DELETE http://localhost:8080/users/<userID>/posts
```


