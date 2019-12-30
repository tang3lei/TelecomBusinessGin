create table packages
(
    id           int           not null identity (1,1),
    name  NVARCHAR(64) not null ,

    desc2        text,
    create_at    datetime      not null DEFAULT CURRENT_TIMESTAMP,
    update_at    datetime      not null DEFAULT CURRENT_TIMESTAMP
        PRIMARY KEY (id)
);
go

create UNIQUE INDEX uniq_package
    on packages (name)
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

