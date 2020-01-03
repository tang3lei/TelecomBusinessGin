create table packages
(
    id           int          not null identity (1,1),
    name         NVARCHAR(64) not null,
    type         int          not null default 0,
    monthly_cost float        not null default 0,
    daily_cost   float        not null default 0,
    desc2        text,
    create_at    datetime     not null DEFAULT CURRENT_TIMESTAMP,
    update_at    datetime     not null DEFAULT CURRENT_TIMESTAMP
        PRIMARY KEY (id)
);
go

create UNIQUE INDEX uniq_package
    on packages (name)
GO

create TRIGGER packages_update_trigger
    on packages
    for update
    as
    update packages
    set update_at = getdate()
    where id in (select DISTINCT id from inserted);
go

ALTER TABLE packages
    ENABLE TRIGGER packages_update_trigger
go

