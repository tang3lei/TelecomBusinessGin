create table deals
(
    id           int           not null identity (1,1),
    deal_name    nvarchar(128) not null default N'未命名',
    deal_time    bigint,
    phone_number VARCHAR(20)   not null,
    job_name   nvarchar(32)  not null,
    type         int,
    cost         float         not null default 0,
    create_at    datetime      not null DEFAULT CURRENT_TIMESTAMP,
    update_at    datetime      not null DEFAULT CURRENT_TIMESTAMP
        PRIMARY KEY (id)
);
go

create UNIQUE INDEX uniq_deal
    on deals (deal_time)
GO

create TRIGGER deals_update_trigger
    on deals
    for update
    as
    update deals
    set update_at = getdate()
    where id in (select DISTINCT id from inserted);
go

ALTER TABLE deals
    ENABLE TRIGGER deals_update_trigger
go

