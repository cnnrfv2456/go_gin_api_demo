CREATE TABLE member(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR (255) NOT NULL,
    account VARCHAR (255) UNIQUE NOT NULL,
    password VARCHAR (255) NOT NULL,
    updated_at timestamp NOT NULL DEFAULT NOW(),
    created_at timestamp NOT NULL DEFAULT NOW()
);

comment on table member is '會員';

comment on column member.name is '名稱';

comment on column member.account is '帳號';

comment on column member.password is '密碼';