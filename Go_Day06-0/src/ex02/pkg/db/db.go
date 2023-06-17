package db

// users
var QueryCreateTableUsers = `
create table if not exists users
(
    id      		bigint primary key generated always as identity,
    login   		varchar(200) not null unique,
    hashed_password varchar(200) not null,
    name    		varchar(200) not null,
    surname 		varchar(200) not null
);`

var QueryAddNewUser = `
insert into users (login, hashed_password, name, surname)
values ($1, $2, $3, $4);
`

var QueryFindUserByLoginAndPassword = `
select id, login, name, surname from users where login = $1 AND hashed_password = $2
`

// posts
var QueryCreateTablePosts = `
create table if not exists posts
(
    id      	bigint primary key generated always as identity,
    userLogin   varchar(200) not null,
    created 	TIMESTAMP DEFAULT NOW(),
    header    	varchar(200) not null,
    content 	varchar(200) not null
);
`

var QueryAddNewPost = `
insert into posts (userLogin, created, header, content)
values ($1, $2, $3, $4);
`

var QueryGetNPosts = `
SELECT userLogin, created, header, content
FROM ( SELECT * FROM posts ORDER BY id DESC LIMIT 3 OFFSET $1) subquery
ORDER BY created DESC;
`

var QueryGetPostsCount = `
SELECT COUNT(*) AS postsCount FROM posts;
`
