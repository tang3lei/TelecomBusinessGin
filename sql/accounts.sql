create table accounts
(
    id           int           not null identity (1,1),
    phone_number VARCHAR(20)   not null,
    user_name    NVARCHAR(20),
    status       int           not null default 0,
    balance      float         not null DEFAULT 0,
    package      NVARCHAR(128) not NULL,
    info         text,
    desc2        text,
    create_at    datetime      not null DEFAULT CURRENT_TIMESTAMP,
    update_at    datetime      not null DEFAULT CURRENT_TIMESTAMP
        PRIMARY KEY (id)
);
go
create UNIQUE INDEX uniq_account
    on accounts (phone_number)
GO

create TRIGGER accounts_update_trigger
    on accounts
    for update
    as
    update accounts
    set update_at = getdate()
    where id in (select DISTINCT id from inserted);
go

ALTER TABLE accounts
    ENABLE TRIGGER accounts_update_trigger
go

